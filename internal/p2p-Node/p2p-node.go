package p2pnode

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"github.com/NetSepio/gateway/contract"
	nodelogs "github.com/NetSepio/gateway/internal/api/handlers/nodes/nodeLogs"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/internal/p2p-Node/host"
	"github.com/NetSepio/gateway/internal/p2p-Node/service"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/load"
)

// DiscoveryInterval is how often we search for other peers via the DHT.
const DiscoveryInterval = time.Second * 10

// DiscoveryServiceTag is used in our DHT advertisements to discover
// other peers.
const DiscoveryServiceTag = "erebrus"

// Node status constants matching the contract's enum
const (
	StatusOffline     uint8 = 0
	StatusOnline      uint8 = 1
	StatusMaintenance uint8 = 2
	StatusDeactivated uint8 = 3
)

// Time thresholds for status changes
const (
	MaintenanceThreshold = 2 * time.Minute
	OfflineThreshold     = 5 * time.Minute
)

// OnlineURI, MaintenanceURI, and OfflineURI are constants for token URIs
const (
	OnlineURI      = "ipfs://bafkreiczwfmevybanlj73w3v2smos2qgoxsfigonmmki4aoftcgike45sq"
	MaintenanceURI = "ipfs://bafybeibil3zpj6povthugmrpwdvhgehrfpbhgkabltrrtwwfijvuguopka"
	OfflineURI     = "ipfs://bafybeicetdyf7ocbdflobb7dkw5lvwejpa6ny3x55ht4pf2cmyedgarxmu"
)

// NodeStateTracker keeps track of node states to minimize contract calls
type NodeStateTracker struct {
	ContractStatus uint8
	LastPing       time.Time
}

// Global map to track node states
var nodeStates = make(map[string]*NodeStateTracker)

func Init() {
	ctx, _ := context.WithCancel(context.Background())

	ha := host.CreateHost()
	ps := service.NewService(ha, ctx)

	bootstrapPeers := []multiaddr.Multiaddr{}
	db := database.GetDB2()

	dht, err := host.NewDHT(ctx, ha, bootstrapPeers)
	if err != nil {
		logrus.Error("failed to init new dht")
		return
	}

	go host.Discover(ctx, ha, dht)

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				var nodes []models.Node
				if err := db.Debug().Model(&models.Node{}).Find(&nodes).Error; err != nil {
					logrus.Error("failed to fetch nodes from db")
					continue
				}

				for _, node := range nodes {
					if _, exists := nodeStates[node.PeerId]; !exists {
						nodeStates[node.PeerId] = &NodeStateTracker{
							ContractStatus: StatusOffline,
							LastPing:       time.Now(),
						}
						continue // Skip first iteration for new nodes
					}

					var (
						newOSInfo     models.OSInfo
						newGeoAddress models.IpGeoAddress
						newIPInfo     models.IPInfo
					)

					// Safely unmarshal JSON data with error handling
					if node.SystemInfo != "" {
						err = json.Unmarshal([]byte(node.SystemInfo), &newOSInfo)
						if err != nil {
							log.Printf("Error unmarshaling newOSInfo from JSON: %v", err)
						}
					}

					if node.IpGeoData != "" && len(node.IpGeoData) > 0 {
						err = json.Unmarshal([]byte(node.IpGeoData), &newGeoAddress)
						if err != nil {
							log.Printf("Error unmarshaling newGeoAddress from JSON: %v", err)
							// Set default values if unmarshal fails
							newGeoAddress = models.IpGeoAddress{
								IpInfoCity:     "Unknown",
								IpInfoCountry:  "Unknown",
								IpInfoLocation: "Unknown",
								IpInfoOrg:      "Unknown",
								IpInfoPostal:   "Unknown",
								IpInfoTimezone: "Unknown",
							}
						}
					} else {
						// Set default values if IpGeoData is empty
						newGeoAddress = models.IpGeoAddress{
							IpInfoCity:     "Unknown",
							IpInfoCountry:  "Unknown",
							IpInfoLocation: "Unknown",
							IpInfoOrg:      "Unknown",
							IpInfoPostal:   "Unknown",
							IpInfoTimezone: "Unknown",
						}
					}

					if node.IpInfo != "" {
						err = json.Unmarshal([]byte(node.IpInfo), &newIPInfo)
						if err != nil {
							log.Printf("Error unmarshaling newIPInfo from JSON: %v", err)
						}
					}

					node.SystemInfo = models.ToJSON(newOSInfo)
					node.IpGeoData = models.ToJSON(newGeoAddress)
					node.IpInfo = models.ToJSON(newIPInfo)

					peerMultiAddr, err := multiaddr.NewMultiaddr(node.PeerAddress)
					if err != nil {
						logrus.Errorf("Invalid peer address for node %s: %v", node.PeerId, err)
						continue
					}

					peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
					if err != nil {
						logrus.Errorf("Failed to get peer info for node %s: %v", node.PeerId, err)
						continue
					}

					isConnected := ha.Connect(ctx, *peerInfo) == nil
					var newStatus uint8
					var nodeStatus string

					if !isConnected {
						timeSinceLastPing := time.Since(nodeStates[node.PeerId].LastPing)
						// Only update status if within our monitoring window
						if timeSinceLastPing <= OfflineThreshold+time.Minute {
							if timeSinceLastPing > OfflineThreshold {
								newStatus = StatusOffline
								nodeStatus = "inactive"
							} else if timeSinceLastPing > MaintenanceThreshold {
								newStatus = StatusMaintenance
								nodeStatus = "inactive"
							} else {
								continue
							}
						} else {
							continue // Skip nodes that have been offline for too long
						}
					} else {
						newStatus = StatusOnline
						nodeStatus = "active"
						nodeStates[node.PeerId].LastPing = time.Now()
					}

					// Update contract status only for peaq nodes
					logrus.Debugf("Chain: %s, Node: %s, status: %s", strings.ToLower(node.Chain), node.PeerId, nodeStatus)
					logrus.Debugf("newStatus: %d, nodeStates[node.PeerId].ContractStatus: %d", newStatus, nodeStates[node.PeerId].ContractStatus)

					if strings.ToLower(node.Chain) == "peaq" && newStatus != nodeStates[node.PeerId].ContractStatus {
						go func(peerId string, status uint8) {
							// Update contract status
							logrus.Info("Updating contract status for node", peerId)
							if err := updateNodeContractStatus(peerId, status); err != nil {
								logrus.Errorf("Failed to update contract status: %v", err)
								return
							}
							nodeStates[peerId].ContractStatus = status
						}(node.PeerId, newStatus)
					}

					// Update database status
					go func(n models.Node, status string) {
						logrus.Info("Updating node status in db for node", n.PeerId)
						n.Status = status
						if status == "active" {
							n.LastPing = time.Now().Unix()
						}
						if err := db.Debug().Save(&n).Error; err != nil {
							logrus.Errorf("Failed to update node: %v", err)
						}
						nodelogs.LogNodeStatus(n.PeerId, status)
					}(node, nodeStatus)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	go service.SubscribeTopics(ps, ha, ctx)
}

// formatNodeId adds the "did:netsepio:" prefix to the peer ID if not present
// func formatNodeId(peerId string) string {
// 	prefix := "did:netsepio:"
// 	if !strings.HasPrefix(peerId, prefix) {
// 		return prefix + peerId
// 	}
// 	return peerId
// }

func updateNodeContractStatus(nodeId string, status uint8) error {
	// Get node data from database to check chain
	db := database.GetDb()
	var node models.Node
	if err := db.Debug().Model(&models.Node{}).Where("peer_id = ?", nodeId).First(&node).Error; err != nil {
		load.Logger.Error("failed to fetch node from db", zap.Error(err))
		return errors.New("failed to fetch node from db: " + err.Error())
	}

	// Only proceed if chain is peaq
	if strings.ToLower(node.Chain) != "peaq" {
		load.Logger.Info("Skipping contract update for non-peaq node", zap.String("nodeId", nodeId), zap.String("chain", node.Chain))
		return nil
	}

	// formattedNodeId := formatNodeId(nodeId)
	formattedNodeId := nodeId

	// Load environment variables if not already loaded
	if os.Getenv("CONTRACT_ADDRESS") == "" {
		err := godotenv.Load()
		if err != nil {
			load.Logger.Error("error loading .env file", zap.Error(err))
			return errors.New("error loading .env file: " + err.Error())
		}
	}

	// Connect to the Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		load.Logger.Error("failed to connect to the Ethereum client", zap.Error(err))
		return errors.New("failed to connect to the Ethereum client: " + err.Error())
	}

	// Create a new instance of the contract
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		load.Logger.Error("failed to instantiate contract", zap.Error(err))
		return errors.New("failed to instantiate contract: " + err.Error())
	}

	// Create auth options for the transaction
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		load.Logger.Error("failed to create private key", zap.Error(err))
		return errors.New("failed to create private key: " + err.Error())
	}

	// Add retry mechanism for getting chain ID
	var chainID *big.Int
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		chainID, err = client.ChainID(context.Background())
		if err == nil {
			break
		}

		if i < maxRetries-1 {
			// Exponential backoff
			delay := time.Duration(1<<uint(i)) * time.Second
			load.Logger.Warn("Failed to get chain ID, retrying", zap.Duration("delay", delay), zap.Error(err))
			time.Sleep(delay)
		}
	}

	if err != nil {
		load.Logger.Error("failed to get chain ID after 5 attempts", zap.Error(err))
		return errors.New("failed to get chain ID after 5 attempts: " + err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		load.Logger.Error("failed to create transactor", zap.Error(err))
		return errors.New("failed to create transactor: " + err.Error())
	}

	// Get node details to fetch tokenId
	opts := &bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}

	// Check if node exists before trying to update
	contractNode, err := instance.Nodes(opts, formattedNodeId)
	if err != nil {
		load.Logger.Error("Failed to get node details from contract", zap.Error(err))
		return errors.New("Failed to get node details from contract: " + err.Error())
	}

	// Check if tokenId is valid
	if contractNode.TokenId == nil || contractNode.TokenId.Cmp(big.NewInt(0)) == 0 {
		load.Logger.Warn("Node exists in database but not in contract or has invalid token ID", zap.String("nodeId", formattedNodeId))
		// Instead of returning error, we'll try to register the node
		return errors.New("Invalid token ID for node " + formattedNodeId)
	}

	// Set gas limit and price
	auth.GasLimit = 300000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		load.Logger.Error("Failed to get gas price", zap.Error(err))
		return errors.New("Failed to get gas price: " + err.Error())
	}
	// Increase gas price by 20% to ensure transaction goes through
	auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	auth.GasPrice = new(big.Int).Div(auth.GasPrice, big.NewInt(100))

	load.Logger.Info("Updating node status", zap.String("nodeId", formattedNodeId), zap.Uint8("status", status))

	// Update node status
	tx, err := instance.UpdateNodeStatus(auth, formattedNodeId, status)
	if err != nil {
		load.Logger.Error("Failed to update node status", zap.Error(err))
		return errors.New("Failed to update node status: " + err.Error())
	}

	load.Logger.Info("Status update transaction sent", zap.String("txHash", tx.Hash().Hex()))

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		load.Logger.Error("Failed to wait for status update transaction", zap.Error(err))
		return errors.New("Failed to wait for status update transaction: " + err.Error())
	}

	if receipt.Status == 0 {
		load.Logger.Error("status update transaction failed", zap.String("nodeId", formattedNodeId))
		return errors.New("status update transaction failed")
	}

	load.Logger.Info("Status update transaction confirmed", zap.String("nodeId", formattedNodeId))

	// Get the appropriate URI based on status
	var uri string
	switch status {
	case StatusOnline:
		uri = OnlineURI
	case StatusMaintenance:
		uri = MaintenanceURI
	case StatusOffline:
		uri = OfflineURI
	default:
		load.Logger.Error("invalid status for URI update", zap.Uint8("status", status))
		return errors.New("invalid status for URI update")
	}

	// Create a new auth for the second transaction
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		load.Logger.Error("failed to create transactor for URI update", zap.Error(err))
		return errors.New("failed to create transactor for URI update: " + err.Error())
	}

	auth.GasLimit = 300000
	auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	auth.GasPrice = new(big.Int).Div(auth.GasPrice, big.NewInt(100))

	load.Logger.Info("Updating token URI", zap.String("nodeId", formattedNodeId), zap.String("uri", uri))

	// Update the token URI
	tx, err = instance.UpdateTokenURI(auth, contractNode.TokenId, uri)
	if err != nil {
		load.Logger.Error("Failed to update token URI", zap.Error(err))
		return errors.New("Failed to update token URI: " + err.Error())
	}

	load.Logger.Info("URI update transaction sent", zap.String("txHash", tx.Hash().Hex()))

	// Wait for transaction to be mined
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		load.Logger.Error("Failed to wait for URI update transaction", zap.Error(err))
		return errors.New("Failed to wait for URI update transaction: " + err.Error())
	}

	if receipt.Status == 0 {
		load.Logger.Error("URI update transaction failed", zap.String("nodeId", formattedNodeId))
		return errors.New("URI update transaction failed")
	}

	load.Logger.Info("Node status and token URI updated",
		zap.String("nodeId", formattedNodeId),
		zap.Uint8("status", status),
		zap.String("uri", uri),
	)
	return nil
}

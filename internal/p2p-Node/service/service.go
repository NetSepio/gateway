package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

const DiscoveryServiceTag = "erebrus"

type status struct {
	Status string
}

func NewService(h host.Host, ctx context.Context) *pubsub.PubSub {
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		logrus.Error(err)
	}
	return ps
}

var Status_data []*Status
var StatusData map[string]*Status

func SubscribeTopics(ps *pubsub.PubSub, h host.Host, ctx context.Context) {
	// Initialize StatusData map
	StatusData = make(map[string]*Status)
	topicString := "status"
	topic, err := ps.Join(DiscoveryServiceTag + "/" + topicString)
	if err != nil {
		logrus.Error(err)
	}
	sub, err := topic.Subscribe()
	if err != nil {
		logrus.Error(err)
	}
	go func() {
		for {
			// Block until we recieve a new message.
			msg, err := sub.Next(ctx)
			if err != nil {
				logrus.Error(err)
				continue
			}
			if msg.ReceivedFrom == h.ID() {
				continue
			}
			var node *models.Node
			if err := json.Unmarshal(msg.Data, &node); err != nil {
				logrus.Error(err)
				continue
			}
			db := database.GetDb()
			node.Status = "active"
			node.LastPing = time.Now().Unix()
			err = CreateOrUpdate(db, node)
			if err != nil {
				logwrapper.Error("failed to update db: ", err.Error())
			}
			if err := topic.Publish(ctx, []byte("Gateway recieved the node information")); err != nil {
				logrus.Error(err)
				continue
			}

			topic.EventHandler()
		}
	}()
	// topicString2 := "client"
	// topic2, err := ps.Join(DiscoveryServiceTag + "/" + topicString2)
	// if err != nil {
	// 	panic(err)
	// }

	// sub2, err := topic2.Subscribe()
	// if err != nil {
	// 	panic(err)
	// }

	// go func() {
	// 	for {
	// 		// Block until we recieve a new message.
	// 		msg, err := sub2.Next(ctx)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if msg.ReceivedFrom == h.ID() {
	// 			continue
	// 		}
	// 		fmt.Printf("[%s] , status isz: %s", msg.ReceivedFrom, string(msg.Data))
	// 		if err := topic2.Publish(ctx, []byte("heres a reply from client")); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }()

}

func CreateOrUpdate(db *gorm.DB, node *models.Node) error {
	var model models.Node

	result := db.Model(&models.Node{}).Where("peer_id = ?", node.PeerId)
	if result.RowsAffected != 0 {
		//exists, update
		return db.Model(&model).Updates(node).Error
	} else {
		// find the list of node which has the node name starting with sg ?
		var nodes []models.Node
		db.Where("peer_id = ?", node.PeerId).Find(&nodes) // find all nodes with name starting with sg
		log.Printf("%+v\n", nodes)
		if len(nodes) > 0 {
			// nodeName, err := bringTopRegionId(nodes, node.Region)
			// if err != nil {
			// 	return err
			// }
			// node.NodeName = nodeName
			return db.Create(node).Error
		} else {
			// if no nodes with name starting with sg, create a new one
			// node.NodeName = node.IpInfoCountry + `001`
			return db.Create(node).Error
		}

	}
	// Use a map to track indices, though it's not necessary for finding the highest number
	// indexMap := make(map[int]models.Node)
	// for i, node := range arr {
	// 	indexMap[i] = node
	// }

	// firstRegionNumber := strings.Split(arr[0].NodeName, region)
	// /*in006 = [,006]*/
	// highest := firstRegionNumber[1]
	// highestInt := 0
	// for _, node := range indexMap {
	// 	splitedNodeName := strings.Split(node.NodeName, region)
	// 	splitedNodeInt, _ := strconv.Atoi(splitedNodeName[1])
	// 	highestInt, _ = strconv.Atoi(highest)
	// 	fmt.Println("splitedNodeInt : ", splitedNodeInt, " , highestInt : ", highestInt)
	// 	if splitedNodeInt > highestInt {
	// 		highestInt = splitedNodeInt
	// 	}
	// }
	// increment the highest number
	// highestInt, err := strconv.Atoi(highest)
	// log.Println("Printing the highest integer { before } = ", highestInt)
	// highestInt++
	// log.Println("Printing the highest integer { after } = ", highestInt)
	// highest = strconv.Itoa(highestInt)
	// for len(highest) < 3 {
	// 	highest = "0" + highest
	// }

	// return region + highest, nil
}

/*func bringTopRegionId(arr []models.Node, region string) (string, error) {
	if len(arr) == 0 {
		return "", fmt.Errorf("array is empty")
	}
	// Use a map to track indices, though it's not necessary for finding the highest number
	indexMap := make(map[int]models.Node)
	for i, node := range arr {
		indexMap[i] = node
	}

	firstRegionNumber := strings.Split(arr[0].NodeName, region)
	// in006 = [,006]
	highest := firstRegionNumber[1]
	highestInt := 0
	for _, node := range indexMap {
		splitedNodeName := strings.Split(node.NodeName, region)
		splitedNodeInt, _ := strconv.Atoi(splitedNodeName[1])
		highestInt, _ = strconv.Atoi(highest)
		if splitedNodeInt > highestInt {
			highestInt = splitedNodeInt
		}
	}
	// increment the highest number
	highestInt++
	highest = strconv.Itoa(highestInt)
	for len(highest) < 3 {
		highest = "0" + highest
	}

	return region + highest, nil
}*/

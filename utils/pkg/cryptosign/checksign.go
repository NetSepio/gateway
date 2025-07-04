package cryptosign

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/minio/blake2b-simd"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrFlowIdNotFound = errors.New("flow id not found")
)

func CheckSign(signature string, flowId string, message string, pubKey string) (string, string, bool, error) {
	db := database.GetDb()
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", "", false, err
	}

	sha3_i := sha3.New256()
	signatureInBytes = append(signatureInBytes, []byte(message)...)
	pubBytes, err := hexutil.Decode(pubKey)
	if err != nil {
		return "", "", false, err
	}
	sha3_i.Write(pubBytes)
	sha3_i.Write([]byte{0})
	hash := sha3_i.Sum(nil)
	addr := hexutil.Encode(hash)

	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", false, ErrFlowIdNotFound
	}

	if !strings.EqualFold(addr, flowIdData.WalletAddress) {
		return "", "", false, err
	}

	msgGot, matches := sign.Open(nil, signatureInBytes, (*[32]byte)(pubBytes))
	if !matches || string(msgGot) != message {
		return "", "", false, err
	}
	return flowIdData.UserId, flowIdData.WalletAddress, true, nil

}

func CheckSignEth(signature string, flowId string, message string) (string, string, bool, error) {

	db := database.GetDb()
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", "", false, err
	}
	if signatureInBytes[64] == 27 || signatureInBytes[64] == 28 {
		signatureInBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		return "", "", false, err
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)
	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", false, ErrFlowIdNotFound
	}
	if err != nil {
		return "", "", false, err
	}
	if strings.EqualFold(flowIdData.WalletAddress, walletAddress.String()) {
		return flowIdData.UserId, flowIdData.WalletAddress, true, nil
	} else {
		return "", "", false, nil
	}
}

func CheckSignSui(signature string, flowId string) (string, string, bool, error) {
	db := database.GetDb()
	// Decode signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return "", "", false, err
	}

	//TODO verify message
	// // Decode message
	// _, err = base64.StdEncoding.DecodeString(message)
	// if err != nil {
	// 	return "", "", false, err
	// }

	// Assuming ED25519 signature format
	size := 32

	publicKey := signatureBytes[len(signatureBytes)-size:]
	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),                       // Curve is not used in serialization
		X:     new(big.Int).SetBytes(publicKey[:]),   // Set X coordinate
		Y:     new(big.Int).SetBytes(publicKey[32:]), // Set Y coordinate
	}
	if pubKey.X == nil || pubKey.Y == nil {
		return "", "", false, err
	}
	// Serialize the public key into bytes
	pubKeyBytes := pubKey.X.Bytes()

	// Pad X coordinate bytes to ensure they are the same length as the curve's bit size
	paddingLen := (pubKey.Curve.Params().BitSize + 7) / 8
	pubKeyBytes = append(make([]byte, paddingLen-len(pubKeyBytes)), pubKeyBytes...)

	// Concatenate the signature scheme flag (0x00 for Ed25519) with the serialized public key bytes
	concatenatedBytes := append([]byte{0x00}, pubKeyBytes...)

	// Compute the BLAKE2b hash
	hash := blake2b.Sum256(concatenatedBytes)

	// The resulting hash is the Sui address
	suiAddress := "0x" + hex.EncodeToString(hash[:])

	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", false, ErrFlowIdNotFound
	}

	if !strings.EqualFold(suiAddress, flowIdData.WalletAddress) {
		return "", "", false, err
	}

	//TODO check from message
	// msgGot, matches := sign.Open(nil, signatureInBytes, (*[32]byte)(pubBytes))
	// if !matches || string(msgGot) != message {
	// 	return "", "", false, err
	// }

	return flowIdData.UserId, flowIdData.WalletAddress, true, nil
}

func CheckSignSol(signature string, flowId string, message string, pubKey string) (string, string, bool, error) {

	db := database.GetDb()
	bytes, err := base58.Decode(pubKey)
	if err != nil {
		return "", "", false, err
	}
	messageAsBytes := []byte(message)

	signedMessageAsBytes, err := hex.DecodeString(signature)

	if err != nil {

		return "", "", false, err
	}

	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", false, err
	}

	ed25519.Verify(bytes, messageAsBytes, signedMessageAsBytes)

	return flowIdData.WalletAddress, flowIdData.UserId, true, nil

}

func VerifySignatureEVM(message, signature, flowId string) (string, string, bool, error) {

	// Decode signature
	sigBytes, err := hex.DecodeString(strings.TrimPrefix(signature, "0x"))
	if err != nil {
		return "", "", false, fmt.Errorf("invalid signature format: %v", err)
	}

	// Ensure signature length is correct
	if len(sigBytes) != 65 {
		return "", "", false, fmt.Errorf("invalid signature length: expected 65, got %d", len(sigBytes))
	}

	// Prepare message hash
	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))

	// Recover public key
	sigBytes[64] -= 27 // Adjust v value for recovery
	pubKey, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		return "", "", false, fmt.Errorf("failed to recover public key: %v", err)
	}
	// Get recovered address
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	var flowIdData models.FlowId
	db := database.GetDb()

	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", false, ErrFlowIdNotFound
	}

	// fmt.Println("Recovered Address:", recoveredAddr.Hex())

	// Compare addresses
	expectedAddr := common.HexToAddress(flowIdData.WalletAddress)
	if recoveredAddr == expectedAddr {
		return flowIdData.UserId, flowIdData.WalletAddress, true, nil
	} else {
		return "", "", false, fmt.Errorf("signature does not match the expected address")
	}
}

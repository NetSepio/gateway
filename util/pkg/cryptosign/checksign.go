package cryptosign

import (
	"errors"
	"fmt"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"crypto/ed25519"
    "encoding/hex"

    "github.com/mr-tron/base58"
)

var (
	ErrFlowIdNotFound = errors.New("flow id not found")
)

func CheckSign(signature string, flowId string, message string, pubKey string) (string, string, bool, error) {
	db := dbconfig.GetDb()
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

	db := dbconfig.GetDb()
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

func CheckSignSol(signature string, flowId string, message string, pubKey string) (string,string, bool, error) {

	db := dbconfig.GetDb()
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
	
	return flowIdData.WalletAddress,flowIdData.UserId,true ,nil

}
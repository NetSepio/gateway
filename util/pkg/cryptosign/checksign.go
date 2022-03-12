package cryptosign

import (
	"errors"
	"fmt"
	"strings"

	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/models"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrFlowIdNotFound = errors.New("flow id not found")
)

func CheckSign(signature string, flowId string, message string) (string, bool, error) {

	db := dbconfig.GetDb()
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", false, err
	}
	if signatureInBytes[64] == 27 || signatureInBytes[64] == 28 {
		signatureInBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		return "", false, err
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)
	var flowIdData models.FlowId
	res := db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData)

	if res.RecordNotFound() {
		return "", false, ErrFlowIdNotFound
	}
	if err := res.Error; err != nil {
		return "", false, err
	}
	if strings.ToLower(flowIdData.WalletAddress) == strings.ToLower(walletAddress.String()) {
		return flowIdData.WalletAddress, true, nil
	} else {
		return "", false, nil
	}
}

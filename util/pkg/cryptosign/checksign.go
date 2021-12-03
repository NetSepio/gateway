package cryptosign

import (
	"errors"
	"fmt"
	"netsepio-api/db"
	"netsepio-api/models"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrFlowIdNotFound = errors.New("Flow id not found")
)

func CheckSign(signature string, flowId string, message string) (string, bool, error) {

	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", false, err
	}
	signatureInBytes[64] -= 27
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		return "", false, err
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)
	var flowIdData models.FlowId
	res := db.Db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData)

	if res.RecordNotFound() {
		return "", false, ErrFlowIdNotFound
	}
	if err := res.Error; err != nil {
		return "", false, err
	}
	fmt.Println("signature", signature)
	fmt.Println(" l ", flowIdData.WalletAddress, " r ", walletAddress.String())
	if flowIdData.WalletAddress == walletAddress.String() {
		return flowIdData.WalletAddress, true, nil
	} else {
		return "", false, nil
	}
}

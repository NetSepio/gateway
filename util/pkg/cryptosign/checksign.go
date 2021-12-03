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
	var user models.User
	res := db.Db.Model(&models.User{}).Where("? = ANY (flow_id)", flowId).First(&user)
	if res.RecordNotFound() {
		return "", false, ErrFlowIdNotFound
	}
	if err != nil {
		return "", false, err
	}

	if user.WalletAddress == walletAddress.String() {
		return user.WalletAddress, true, nil
	} else {
		return "", false, nil
	}
}

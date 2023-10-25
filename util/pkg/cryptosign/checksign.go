package cryptosign

import (
	"errors"
	"log"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ErrFlowIdNotFound = errors.New("flow id not found")
)

func CheckSign(signature string, flowId string, message string, pubKey string) (string, bool, error) {

	db := dbconfig.GetDb()
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", false, err
	}

	sha3_i := sha3.New256()
	signatureInBytes = append(signatureInBytes, []byte(message)...)
	pubBytes, err := hexutil.Decode(pubKey)
	if err != nil {
		return "", false, err
	}
	sha3_i.Write(pubBytes)
	sha3_i.Write([]byte{0})
	hash := sha3_i.Sum(nil)
	addr := hexutil.Encode(hash)
	log.Printf("pub key - %v\n", hexutil.Encode(pubBytes))

	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", false, ErrFlowIdNotFound
	}
	if addr != flowIdData.WalletAddress {
		return "", false, err
	}
	msgPro := signatureInBytes[sign.Overhead:]
	log.Printf("msg pro - %v\n", string(msgPro))

	msgGot, matches := sign.Open(nil, signatureInBytes, (*[32]byte)(pubBytes))
	log.Printf("msg got - %v\n", string(msgGot))
	log.Printf("msg needed - %v\n", message)
	if !matches || string(msgGot) != message {
		log.Println("no match or no equal")
		return "", false, err
	}
	return flowIdData.WalletAddress, true, nil

}

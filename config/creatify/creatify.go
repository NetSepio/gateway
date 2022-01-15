package creatify

import (
	"errors"

	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
)

type tRole int

var (
	ErrRoleNotExist = errors.New("role does not exist")
)

const (
	CREATOR_ROLE  tRole = iota
	ADMIN_ROLE    tRole = iota
	OPERATOR_ROLE tRole = iota
)

type tRoles map[tRole][32]byte

var roles tRoles = tRoles{}
var initiated = false

func GetRole(role tRole) ([32]byte, error) {
	if !initiated {
		InitRolesId()
	}
	v, ok := roles[role]
	if !ok {
		return [32]byte{}, ErrRoleNotExist
	}
	return v, nil
}
func InitRolesId() {

	instance := GetInstance(smartcontract.GetClient())
	creatorRoleId, err := instance.CREATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "CREATORROLE", err.Error())
	}
	roles[CREATOR_ROLE] = creatorRoleId
	adminRoleId, err := instance.DEFAULTADMINROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "DEFAULTADMINROLE", err.Error())
	}
	roles[ADMIN_ROLE] = adminRoleId

	operatorRoleId, err := instance.OPERATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "OPERATORROLE", err.Error())
	}
	roles[OPERATOR_ROLE] = operatorRoleId
}

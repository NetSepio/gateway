package netsepio

import (
	"errors"

	"github.com/NetSepio/gateway/config/smartcontract"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
)

type tRole int

var (
	ErrRoleNotExist = errors.New("role does not exist")
)

const (
	VOTER_ROLE    tRole = iota
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
	client, err := smartcontract.GetClient()
	if err != nil {
		logwrapper.Fatalf("failed to client, error: %v", err.Error())
	}
	instance, err := GetInstance(client)
	if err != nil {
		logwrapper.Fatalf("failed to get instance for %v , error: %v", "NETSEPIO", err.Error())
	}
	voterRoleId, err := instance.NETSEPIOVOTERROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "NETSEPIOVOTERROLE", err.Error())
	}
	roles[VOTER_ROLE] = voterRoleId
	adminRoleId, err := instance.NETSEPIOADMINROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "NETSEPIOADMINROLE", err.Error())
	}
	roles[ADMIN_ROLE] = adminRoleId

	operatorRoleId, err := instance.NETSEPIOMODERATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "NETSEPIOMODERATORROLE", err.Error())
	}
	roles[OPERATOR_ROLE] = operatorRoleId
}

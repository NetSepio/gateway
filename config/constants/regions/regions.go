package regions

import "github.com/NetSepio/gateway/config/envconfig"

type Region struct {
	Name       string
	Code       string
	ServerHttp string
}

var Regions map[string]Region
var ErebrusRegions map[string]Region

func InitRegions() {
	Regions = map[string]Region{
		"us-east-2": {
			Name:       "US east 2",
			Code:       "us-east-2",
			ServerHttp: envconfig.EnvVars.VPN_DEPLOYER_API_US_EAST,
		},
		"ap-southeast-1": {
			Name:       "Asia Pacific",
			Code:       "ap-southeast-1",
			ServerHttp: envconfig.EnvVars.VPN_DEPLOYER_API_SG,
		},
	}

	ErebrusRegions = map[string]Region{
		"us-east-2": {
			Name:       "US east 2",
			Code:       "us-east-2",
			ServerHttp: envconfig.EnvVars.EREBRUS_API_US_EAST,
		},
		"ap-southeast-1": {
			Name:       "Asia Pacific",
			Code:       "ap-southeast-1",
			ServerHttp: envconfig.EnvVars.EREBRUS_API_SG,
		},
	}
}

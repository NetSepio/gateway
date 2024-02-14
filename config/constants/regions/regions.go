package regions

import "github.com/NetSepio/gateway/config/envconfig"

type Region struct {
	Name       string
	Code       string
	ServerHttp string
}

var Regions map[string]Region
var ErebrusRegions map[string]Region

//todo
//Store regions in persistent DB
//

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
		"us": {
			Name:       "US",
			Code:       "us",
			ServerHttp: envconfig.EnvVars.SOTREUS_US,
		},
		"sg": {
			Name:       "Singapore",
			Code:       "sg",
			ServerHttp: envconfig.EnvVars.SOTREUS_SG,
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
		"us": {
			Name:       "US",
			Code:       "us",
			ServerHttp: envconfig.EnvVars.EREBRUS_US,
		},
		"sg": {
			Name:       "Singapore",
			Code:       "sg",
			ServerHttp: envconfig.EnvVars.EREBRUS_SG,
		},
		"ca": {
			Name:       "Canada",
			Code:       "ca",
			ServerHttp: envconfig.EnvVars.EREBRUS_CA,
		},
		"eu": {
			Name:       "Europe",
			Code:       "eu",
			ServerHttp: envconfig.EnvVars.EREBRUS_EU,
		},
		"jp": {
			Name:       "Japan",
			Code:       "jp",
			ServerHttp: envconfig.EnvVars.EREBRUS_JP,
		},
		"hk": {
			Name:       "Hong Kong",
			Code:       "hk",
			ServerHttp: envconfig.EnvVars.EREBRUS_HK,
		},
		"hk02": {
			Name:       "Hong Kong 02",
			Code:       "hk02",
			ServerHttp: envconfig.EnvVars.EREBRUS_HK_02,
		},
	}
}

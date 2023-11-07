package envconfig

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	PASETO_PRIVATE_KEY        string        `env:"PASETO_PRIVATE_KEY,required"`
	PASETO_EXPIRATION         time.Duration `env:"PASETO_EXPIRATION,required"`
	APP_PORT                  int           `env:"APP_PORT,required"`
	AUTH_EULA                 string        `env:"AUTH_EULA,required"`
	VOTER_EULA                string        `env:"VOTER_EULA,required"`
	APP_NAME                  string        `env:"APP_NAME,required"`
	GIN_MODE                  string        `env:"GIN_MODE,required"`
	DB_HOST                   string        `env:"DB_HOST,required"`
	DB_USERNAME               string        `env:"DB_USERNAME,required"`
	DB_PASSWORD               string        `env:"DB_PASSWORD,required"`
	DB_NAME                   string        `env:"DB_NAME,required"`
	DB_PORT                   int           `env:"DB_PORT,required"`
	NETSEPIO_CONTRACT_ADDRESS string        `env:"NETSEPIO_CONTRACT_ADDRESS,required"`
	POLYGON_RPC               string        `env:"POLYGON_RPC,required"`
	MNEMONIC                  string        `env:"MNEMONIC,required"`
	IPFS_NODE_URL             string        `env:"IPFS_NODE_URL,required"`
	ALLOWED_ORIGIN            []string      `env:"ALLOWED_ORIGIN,required" envSeparator:","`
	GRAPH_API                 string        `env:"GRAPH_API,required"`
	SIGNED_BY                 string        `env:"SIGNED_BY,required"`
	FUNCTION_ID               string        `env:"FUNCTION_ID,required"`
	GAS_UNITS                 int           `env:"GAS_UNITS,required"`
	GAS_PRICE                 int           `env:"GAS_PRICE,required"`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}

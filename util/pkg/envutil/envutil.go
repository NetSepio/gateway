package envutil

import (
	"os"

	"github.com/NetSepio/gateway/util/pkg/logwrapper"
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logwrapper.Fatalf("env variable %v is not defined", key)
	}
	return val
}

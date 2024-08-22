package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/NetSepio/gateway/app"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
)

func main() {
	app.Init()

	db := dbconfig.GetDb()
	// pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	// fmt.Printf("priv = %s\npub = %s\n", hex.EncodeToString(priv), hex.EncodeToString(pub))
	if os.Getenv("DEBUG_MODE") == "true" {
		walletAddrLower := strings.ToLower("0xdd3933022e36e9a0a15d0522e20b7b580d38b54ec9cb28ae09697ce0f7c95b6b")
		newUser := &models.User{
			WalletAddress: &walletAddrLower,
			UserId:        "fc8fe270-ce16-4df9-a17f-979bcd824e32",
		}
		if err := db.Create(newUser).Error; err != nil {
			logwrapper.Warn(err)
		}
		newClaims := claims.NewWithWallet(newUser.UserId, newUser.WalletAddress)

		pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			panic(err)
		}
		token, err := auth.GenerateToken(newClaims, ed25519.PrivateKey(pvKey))
		if err != nil {
			panic(err)
		}
		fmt.Printf("========TEST TOKEN========\n%s\n========TEST TOKEN========\n", token)
	}
	dbconfig.Init()
	logwrapper.Log.Info("Starting app")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}

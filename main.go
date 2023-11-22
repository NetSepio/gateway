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
		newUser := &models.User{
			WalletAddress: strings.ToLower("0x984185d39c67c954bd058beb619faf8929bb9349ef33c15102bdb982cbf7f18f"),
		}
		if err := db.Create(newUser).Error; err != nil {
			logwrapper.Warn(err)

		}
		newClaims := claims.New(newUser.WalletAddress)

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
	logwrapper.Log.Info("Starting app")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}

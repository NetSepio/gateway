package routines

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/robfig/cron"
	"netsepio-gateway-v1.1/internal/api/handlers/leaderboard"
	"netsepio-gateway-v1.1/internal/api/handlers/referral"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/models/claims"
	"netsepio-gateway-v1.1/utils/auth"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func AutoCalculateScoreBoard() {

	db := database.GetDb()
	if load.Cfg.GIN_MODE == strings.ToLower("debug") {
		walletAddrLower := strings.ToLower("0xdd3933022e36e9a0a15d0522e20b7b580d38b54ec9cb28ae09697ce0f7c95b6b")
		// first check if the user already exists
		var user models.User
		if err := db.Where("wallet_address = ?", walletAddrLower).First(&user).Error; err != nil {
			logwrapper.Warn(err)
		}
		if user.UserId != "" {
			logwrapper.Warn("User already exists")
			return
		}
		newUser := &models.User{
			WalletAddress: &walletAddrLower,
			UserId:        "fc8fe270-ce16-4df9-a17f-979bcd824e32",
			ReferralCode:  referral.GetReferalCode(),
		}
		if err := db.Create(newUser).Error; err != nil {
			logwrapper.Warn(err)
		}
		newClaims := claims.NewWithWallet(newUser.UserId, newUser.WalletAddress)

		pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			panic(err)
		}
		token, err := auth.GenerateToken(newClaims, ed25519.PrivateKey(pvKey))
		if err != nil {
			panic(err)
		}
		fmt.Printf("========TEST TOKEN========\n%s\n========TEST TOKEN========\n", token)
	}

	go func() {
		c := cron.New()
		// Schedule the function to run every day at midnight (or adjust the schedule as needed)
		c.AddFunc("0 30 18 * * *", func() {
			leaderboard.AutoCalculateScoreBoard()
		})

		// Start the cron scheduler in the background
		c.Start()

		// Keep the application running
		select {}
	}()
}

func Init() {
	AutoCalculateScoreBoard()
}

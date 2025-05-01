package email

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/caching"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/load"
)

var (
	rdb = caching.Rdb
	ctx = context.Background() // context.Background() is a function that returns a new context
)

func generateOTP() string {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%06s", n.String())
}

func SendOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	otp := generateOTP()
	expiration := 15 * time.Minute

	rdb.Set(ctx, otp, req.Email, expiration)

	client := resend.NewClient(load.Cfg.RESEND_API_KEY)
	params := &resend.SendEmailRequest{
		To:      []string{req.Email},
		From:    "Erebrus Info <noreply@info.erebrus.io>",
		Text:    fmt.Sprintf("Your OTP for verifying this email is: %s", otp),
		Subject: "Your OTP Code",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func VerifyOTP(c *gin.Context) {
	var req struct {
		OTP string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	storedEmail, err := rdb.Get(ctx, req.OTP).Result()
	if err == redis.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	} else {

		db := database.GetDb()

		userId := c.GetString(paseto.CTX_USER_ID)

		// print the value of the key
		logrus.Infoln("storedEmail : ", storedEmail)

		// update user's email in the database
		err := db.Model(&models.User{}).Where("user_id = ?", userId).Update("email", storedEmail).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "message": "failed to update the details please try again"})
				return

			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update email"})
			return
		}
	}

	rdb.Del(ctx, req.OTP) // OTP is one-time use

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

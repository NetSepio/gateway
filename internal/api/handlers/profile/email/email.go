package email

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/caching"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
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
	fmt.Println("OTP : ", otp)
	expiration := 15 * time.Minute

	status := caching.Rdb.Set(ctx, otp, req.Email, expiration)
	if status.Err() != nil {
		logrus.Errorf("failed to set OTP in redis: %s", status.Err())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache OTP"})
		return
	}

	client := resend.NewClient(load.Cfg.RESEND_API_KEY)
	params := &resend.SendEmailRequest{
		From:    "noreply@info.erebrus.io", // Must be verified
		To:      []string{req.Email},
		Subject: "Your OTP Code",
		Text:    fmt.Sprintf("Your OTP for verifying this email is: %s", otp),
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		logrus.Infof("failed to send OTP via resend: %s", err)
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

	storedEmail, err := caching.Rdb.Get(ctx, req.OTP).Result()
	if err != nil && len(storedEmail) == 0 {
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
		} else {
			caching.Rdb.Del(ctx, req.OTP) // OTP is one-time use

			c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
			return

		}
	}

}

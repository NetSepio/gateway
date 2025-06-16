package claims

import (
	"errors"
	"fmt"
	"time"

	"github.com/vk-rv/pvx"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/load"
)

type CustomClaims struct {
	WalletAddress *string `json:"walletAddress,omitempty"`
	UserId        string  `json:"userId"`
	SignedBy      string  `json:"signedBy"`
	Email         *string `json:"email,omitempty"`
	pvx.RegisteredClaims
}

type CustomClaimsForOrganisation struct {
	OrganisationName *string `json:"orgainisationName,omitempty"`
	OrganisationId   string  `json:"organisationId,omitempty"`
	SignedBy         string  `json:"signedBy"`
	IpAddress        *string `json:"ipAddress"`
	pvx.RegisteredClaims
}

// Valid implements pvx.Claims.
func (c CustomClaimsForOrganisation) Valid() error {
	db := database.GetDb()
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	fmt.Printf("c.OrganisationId: %s\n", c.OrganisationId)
	err := db.Model(&models.Organisation{}).Where("id = ?", c.OrganisationId).First(&models.Organisation{}).Error
	if err != nil {
		return err
	}
	return nil
}

type AuthClaim struct {
	AuthId string `json:"authId"`
	pvx.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	db := database.GetDb()
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	fmt.Printf("c.UserId: %s\n", c.UserId)
	if len(c.UserId) == 0 {
		return errors.New("user id is empty")
	} else {
		err := db.Model(&models.User{}).Where("user_id = ?", c.UserId).First(&models.User{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func NewWithWallet(userId string, walletAddr *string) CustomClaims {
	pasetoExpirationInHours := load.Cfg.PASETO_EXPIRATION
	expiration := time.Now().Add(pasetoExpirationInHours)
	signedBy := load.Cfg.PASETO_SIGNED_BY
	return CustomClaims{
		walletAddr,
		userId,
		signedBy,
		nil,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}

func NewWithOrganisation(organisationId string, organisationName *string, ipAddress *string) CustomClaimsForOrganisation {
	pasetoExpirationInHours := load.Cfg.PASETO_EXPIRATION
	expiration := time.Now().Add(pasetoExpirationInHours)
	signedBy := load.Cfg.PASETO_SIGNED_BY
	return CustomClaimsForOrganisation{
		organisationName,
		organisationId,
		signedBy,
		ipAddress,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}

func (c AuthClaim) Valid() error {
	// check if authId exists in db and is not expired by 5 minutes
	db := database.GetDb()
	var emailAuth models.EmailAuth
	err := db.Model(&models.EmailAuth{}).Where("id = ?", c.AuthId).First(&emailAuth).Error
	if err != nil {
		return err
	}

	if time.Since(emailAuth.CreatedAt) > 5*time.Minute {
		return fmt.Errorf("auth id expired")
	}
	return nil
}

func NewWithEmail(userId string, email *string) CustomClaims {
	pasetoExpirationInHours := load.Cfg.PASETO_EXPIRATION
	expiration := time.Now().Add(pasetoExpirationInHours)
	signedBy := load.Cfg.PASETO_SIGNED_BY
	return CustomClaims{
		nil,
		userId,
		signedBy,
		email,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}

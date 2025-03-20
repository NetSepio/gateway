package referral

import (
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
)

// CreateReferralDiscount inserts a new referral discount into the database
func CreateReferralDiscount(referral *models.ReferralDiscount) error {
	db := dbconfig.GetDb()
	return db.Create(referral).Error
}

// GetReferralDiscount retrieves a referral discount by ID
func GetReferralDiscount(id string) (referral models.ReferralDiscount, err error) {
	db := dbconfig.GetDb()
	err = db.Where("id = ?", id).First(&referral).Error
	return referral, err
}

func GetReferralDiscountByUserId(user_id string) (referral models.ReferralDiscount, err error) {
	db := dbconfig.GetDb()
	err = db.Where("user_id = ?", user_id).First(&referral).Error
	return referral, err
}

func GetReferralDiscountByReferralCode(referral_code string) (referral models.ReferralDiscount, err error) {
	db := dbconfig.GetDb()
	err = db.Where("referral_code = ?", referral_code).First(&referral).Error
	return referral, err
}

// GetAllReferralDiscounts retrieves all referral discounts
func GetAllReferralDiscounts() (referrals []models.ReferralDiscount, err error) {
	db := dbconfig.GetDb()
	err = db.Find(&referrals).Error
	return referrals, err
}

// UpdateReferralDiscount updates an existing referral discount
func UpdateReferralDiscount(referral *models.ReferralDiscount) error {
	db := dbconfig.GetDb()
	return db.Save(referral).Error
}

// DeleteReferralDiscount deletes a referral discount by ID
func DeleteReferralDiscount(id string) error {
	db := dbconfig.GetDb()
	return db.Where("id = ?", id).Delete(&models.ReferralDiscount{}).Error
}

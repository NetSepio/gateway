package useractivity

import (
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/logwrapper"
)

func Save(u models.UserActivity) {

	db := database.GetDb()

	if err := db.Create(&u).Error; err != nil {
		logwrapper.Errorf("UserActivity Failed to save this error : %+v\n", err)
		return
	}
}

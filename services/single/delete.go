package single

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func DeleteOne(userID uint) error {
	db := database.GetDB()

	result := db.Model(&model.User{}).Delete("id = ?", userID)
	if result.Error != nil {
		return fmt.Errorf("%s", "delete failed")
	}

	return nil
}

func DeleteMany(userIDs *[]uint) error {
	db := database.GetDB()

	result := db.Model(&model.User{}).Delete("id in ?", *userIDs)
	if result.Error != nil {
		return fmt.Errorf("%s", "delete failed")
	}

	return nil
}

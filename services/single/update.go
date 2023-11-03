package single

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
	"gorm.io/gorm"
)

func UpdateOne(userID uint) error {
	db := database.GetDB()

	result := db.Model(&model.User{}).Where("id = ?", userID).Update("age", gorm.Expr("age + ?", 1))
	if result.Error != nil {
		return fmt.Errorf("%s", "update failed")
	}

	return nil
}

func UpdateMany(userIDs *[]uint) error {
	db := database.GetDB()

	result := db.Model(&model.User{}).Where("id in ?", *userIDs).Update("born", gorm.Expr("DATE_ADD(born, interval 1 second)"))
	if result.Error != nil {
		return fmt.Errorf("%s", "update failed")
	}

	return nil
}

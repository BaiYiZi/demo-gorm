package single

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func CreateOne(user *model.User) (uint, error) {
	db := database.GetDB()

	result := db.Model(&model.User{}).Create(user)
	if result.Error != nil {
		return 0, fmt.Errorf("%s", "create failed")
	}

	return user.ID, nil
}

func CreateMany(userItem *[]model.User) (*[]uint, error) {
	db := database.GetDB()

	result := db.Model(&model.User{}).Create(userItem)
	if result.Error != nil {
		return nil, fmt.Errorf("%s", "create failed")
	}

	insertUserIDs := []uint{}
	for _, v := range *userItem {
		insertUserIDs = append(insertUserIDs, v.ID)
	}

	return &insertUserIDs, nil
}

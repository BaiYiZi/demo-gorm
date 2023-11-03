package single

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func RetrieveOne(userID uint) (*model.User, error) {
	db := database.GetDB()

	resultUser := &model.User{}
	result := db.Model(&model.User{}).Where("id = ?", userID).Find(resultUser)
	if result.Error != nil {
		return nil, fmt.Errorf("%s", "query failed")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return resultUser, nil
}

func RetrieveMany(IDs *[]uint) (*[]model.User, error) {
	db := database.GetDB()

	resultUsers := &[]model.User{}
	result := db.Model(&model.User{}).Where("id in ?", *IDs).Find(resultUsers)
	if result.Error != nil {
		return nil, fmt.Errorf("%s", "query failed")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return resultUsers, nil
}

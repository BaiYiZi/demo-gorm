package relation

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func CreateBoyAndGirl(boy *model.Boy, girl *model.Girl) (*model.Boy, *model.Girl, error) {
	db := database.GetDB()

	if err := db.Model(&model.Boy{}).Omit("GirlFriendID").Create(boy).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create boy failed")
	}

	if err := db.Model(&model.Girl{}).Create(girl).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create girl failed")
	}

	return boy, girl, nil
}

func ConfirmBoyAndGirlRelation(boyID uint, girlID uint) error {
	db := database.GetDB()

	if err := db.Model(&model.Boy{}).Where("id = ?", boyID).Update("GirlFriendID", girlID).Error; err != nil {
		return fmt.Errorf("confirm relation failed")
	}

	return nil
}

func CancelBoyAndGirlRelation(boyID uint) error {
	db := database.GetDB()

	if err := db.Model(&model.Boy{}).Where("id = ?", boyID).Update("GirlFriendID", nil).Error; err != nil {
		return fmt.Errorf("cancel relation failed")
	}

	return nil
}

func SelectGirlByBoy(boyID uint) (any, error) {
	db := database.GetDB()

	queryStruct := struct {
		BoyID           uint   `json:"boy_id" gorm:"column:boy_id"`
		BoyName         string `json:"boy_name" gorm:"column:boy_name"`
		BoyAge          int    `json:"boy_age" gorm:"column:boy_age"`
		BoyGirlFriendID uint   `json:"boy_girl_friend_id" gorm:"column:boy_girl_friend_id"`
		GirlID          uint   `json:"girl_id" gorm:"column:girl_id"`
		GirlName        string `json:"girl_name" gorm:"column:girl_name"`
		GirlAge         int    `json:"girl_age" gorm:"column:girl_age"`
	}{}

	result := db.
		Table(model.Boy{}.TableName()+" boy").
		Select(
			"boy.id as boy_id",
			"boy.name as boy_name",
			"boy.age as boy_age",
			"boy.girl_friend_id as boy_girl_friend_id",
			"girl.id as girl_id",
			"girl.name as girl_name",
			"girl.age as girl_age",
		).
		InnerJoins(
			"JOIN table_girl girl ON boy.girl_friend_id = girl.id WHERE " + fmt.Sprintf("boy.id = %d", boyID),
		).Find(&queryStruct)
	if result.Error != nil {
		return nil, fmt.Errorf("query failed")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &queryStruct, nil
}

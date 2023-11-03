package relation

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func CreatePeopleAndPhone(people *model.People, phone *model.Phone) (*model.People, *model.Phone, error) {
	db := database.GetDB()

	if err := db.Model(&model.People{}).Create(people).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create people failed")
	}

	if err := db.Model(&model.Phone{}).Omit("PeopleID").Create(phone).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create phone failed")
	}

	return people, phone, nil
}

func ConfirmPeopleAndPhoneRelation(peopleID uint, phoneIDs *[]uint) error {
	db := database.GetDB()

	if err := db.Model(&model.Phone{}).Where("id in ?", *phoneIDs).Update("PeopleID", peopleID).Error; err != nil {
		return fmt.Errorf("confirm relation failed")
	}

	return nil
}

func CancelPeopleAndPhoneRelation(peopleID uint, phoneIDs *[]uint) error {
	db := database.GetDB()

	if err := db.Model(&model.Phone{}).Where("id in ?", *phoneIDs).Update("PeopleID", nil).Error; err != nil {
		return fmt.Errorf("confirm relation failed")
	}

	return nil
}

func SelectPhoneByPeople(peopleID uint) (any, error) {
	db := database.GetDB()

	queryStruct := []struct {
		PeopleID      uint   `json:"people_id" gorm:"column:people_id"`
		PeopleName    string `json:"people_name" gorm:"column:people_name"`
		PhoneID       int    `json:"phone_id" gorm:"column:phone_id"`
		PhonePeopleID uint   `json:"phone_people_id" gorm:"column:phone_people_id"`
		PhoneCode     uint   `json:"phone_code" gorm:"column:phone_code"`
	}{}

	result := db.
		Table(model.People{}.TableName()+" people").
		Select(
			"people.id as people_id",
			"people.name as people_name",
			"phone.id as phone_id",
			"phone.people_id as phone_people_id",
			"phone.code as phone_code",
		).
		InnerJoins(
			"JOIN table_phone phone ON people.id = phone.people_id WHERE " + fmt.Sprintf("people.id = %d", peopleID),
		).Find(&queryStruct)
	if result.Error != nil {
		return nil, fmt.Errorf("query failed")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &queryStruct, nil
}

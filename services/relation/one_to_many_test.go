package relation

import (
	"testing"

	"github.com/BaiYiZi/learn-gorm/model"
)

func TestCreatePeopleAndPhone(t *testing.T) {
	people := &model.People{
		Name: "black cat",
	}

	phone := &model.Phone{
		Code: "12345678987",
	}

	_, _, err := CreatePeopleAndPhone(people, phone)
	if err != nil {
		t.Error(err)
	}

	t.Logf("people id %d, phone id %d", people.ID, phone.ID)
}

func TestConfirmPeopleAndPhoneRelation(t *testing.T) {
	var peopleID uint = 1
	phoneIDs := &[]uint{1, 2, 3}

	err := ConfirmPeopleAndPhoneRelation(peopleID, phoneIDs)
	if err != nil {
		t.Error(t)
	}

	t.Logf("confirm relation success people id %d and phone ids %v", peopleID, phoneIDs)
}

func TestCancelPeopleAndPhoneRelation(t *testing.T) {
	var peopleID uint = 1
	phoneIDs := &[]uint{1, 2, 3}

	err := CancelPeopleAndPhoneRelation(peopleID, phoneIDs)
	if err != nil {
		t.Error(t)
	}

	t.Logf("cancel relation success people id %d and phone ids %v", peopleID, phoneIDs)
}

func TestSelectPhoneByPeople(t *testing.T) {
	var peopleID uint = 1

	res, err := SelectPhoneByPeople(peopleID)
	if err != nil {
		t.Error(t)
	}

	t.Log(res)
}

package relation

import (
	"testing"

	"github.com/BaiYiZi/learn-gorm/model"
)

func TestCreateBoyAndGirl(t *testing.T) {
	boy := &model.Boy{
		Name: "boy1",
		Age:  20,
	}

	girl := &model.Girl{
		Name: "girl1",
		Age:  20,
	}

	_, _, err := CreateBoyAndGirl(boy, girl)
	if err != nil {
		t.Error(err)
	}

	t.Logf("boy id %d, girl id %d", boy.ID, girl.ID)
}

func TestConfirmBoyAndGirlRelation(t *testing.T) {
	var boyID uint = 1
	var girlID uint = 3

	err := ConfirmBoyAndGirlRelation(boyID, girlID)
	if err != nil {
		t.Error(err)
	}

	t.Logf("confirm relation success boy id %d and girl id %d", boyID, girlID)
}

func TestCancelBoyAndGirlRelation(t *testing.T) {
	var boyID uint = 1

	err := CancelBoyAndGirlRelation(boyID)
	if err != nil {
		t.Error(err)
	}

	t.Log("cancel relation success boy id", boyID)
}

func TestSelectGirlByBoy(t *testing.T) {
	var boyID uint = 1

	res, err := SelectGirlByBoy(boyID)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

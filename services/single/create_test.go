package single

import (
	"testing"
	"time"

	"github.com/BaiYiZi/learn-gorm/model"
)

func TestCreateOne(t *testing.T) {
	user := &model.User{
		Name:    "black cat",
		Age:     20,
		Address: "二环外的小区",
		Born:    time.Now().UTC(),
	}

	insertID, err := CreateOne(user)

	if err != nil {
		t.Error(err)
	}

	t.Log("insert user id", insertID)
}

func TestCreateMany(t *testing.T) {
	users := &[]model.User{
		{Name: "black cat", Age: 26, Address: "警察局", Born: time.Now().UTC()},
		{Name: "white snow", Age: 18, Address: "城堡", Born: time.Now().UTC()},
		{Name: "black mountain", Age: 120, Address: "cave", Born: time.Now().UTC()},
	}

	insertIDs, err := CreateMany(users)

	if err != nil {
		t.Error(err)
	}

	t.Log("insert user ids", insertIDs)
}

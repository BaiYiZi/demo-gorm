package single

import "testing"

func TestRetrieveOne(t *testing.T) {
	var id uint = 11

	user, err := RetrieveOne(id)
	if err != nil {
		t.Error(err)
	}

	t.Log(user)
}

func TestRetrieveMany(t *testing.T) {
	ids := &[]uint{10, 11}

	users, err := RetrieveMany(ids)
	if err != nil {
		t.Error(err)
	}

	t.Log(users)
}

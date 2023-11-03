package single

import "testing"

func TestDeleteOne(t *testing.T) {
	var delID uint = 2

	err := DeleteOne(delID)
	if err != nil {
		t.Error(err)
	}

	t.Log("delete user id", delID)
}

func TestDeleteMany(t *testing.T) {
	ids := &[]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	err := DeleteMany(ids)
	if err != nil {
		t.Error(err)
	}

	t.Log("delete user ids", ids)
}

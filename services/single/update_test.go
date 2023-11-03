package single

import "testing"

func TestUpdateOne(t *testing.T) {
	var updateID uint = 11

	err := UpdateOne(updateID)
	if err != nil {
		t.Error(err)
	}

	t.Log("update id", updateID)
}

func TestUpdateMany(t *testing.T) {
	updateIDs := &[]uint{10, 11}

	err := UpdateMany(updateIDs)
	if err != nil {
		t.Error(err)
	}

	t.Log("update ids", updateIDs)
}

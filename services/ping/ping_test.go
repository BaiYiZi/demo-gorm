package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
	res, err := Ping("baiyizi")

	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

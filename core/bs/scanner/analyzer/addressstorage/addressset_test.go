package addressstorage

import (
	"testing"
)

func Test_Set(t *testing.T) {

	db := NewSet()

	db.SetValue("hi")
	b := db.HasValue("hi")

	if b != true {
		t.Fatal("db get fail ")
	}
}

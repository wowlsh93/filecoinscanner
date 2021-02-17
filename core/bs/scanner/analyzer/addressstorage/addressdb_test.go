package addressstorage

import (
	"testing"
)

func Test_DB(t *testing.T) {

	db2, err2 := NewDB("./test/")

	if err2 != nil {
		t.Fatal("db get fail ")
	}

	db2.SetValue("hi")
	b := db2.HasValue("hi")

	if b != true {
		t.Fatal("db get fail ")
	}

}

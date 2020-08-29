package db_test

import (
	"testing"

	"github.com/srobles-globant/golang-academy-db/db"
)

func TestConnection(t *testing.T) {
	cases := []struct {
		name  string
		dbImp db.Db
	}{
		{name: "connection to InMemoryDB", dbImp: &db.InMemoryDb{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if res := c.dbImp.Connected(); res {
				t.Errorf("expected %t but got %t", false, res)
			}
			if ok := c.dbImp.Connect(); !ok {
				t.Errorf("expected %t but got %t", true, ok)
			}
		})
	}
}

func TestCreate(t *testing.T) {}

func TestRetrieve(t *testing.T) {}

func TestUpdate(t *testing.T) {}

func TestDelete(t *testing.T) {}

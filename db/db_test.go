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
			t.Run("start not connected", func(t *testing.T) {
				if state := c.dbImp.Connected(); state {
					t.Errorf("expected connected state to be %t but got %t", false, state)
				}
			})
			t.Run("connect operation", func(t *testing.T) {
				if ok := c.dbImp.Connect(); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
			})
			t.Run("connected state", func(t *testing.T) {
				if state := c.dbImp.Connected(); !state {
					t.Errorf("expected connected state to be %t but got %t", true, state)
				}
			})
			t.Run("diconnected state", func(t *testing.T) {
				c.dbImp.Disconnect()
				if state := c.dbImp.Connected(); state {
					t.Errorf("expected connected state to be %t but got %t", false, state)
				}
			})
		})
	}
}

func TestCRUD(t *testing.T) {
	cases := []struct {
		name  string
		dbImp db.Db
	}{
		{name: "CRUD operations with InMemoryDB", dbImp: &db.InMemoryDb{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.dbImp.Connect()
			var id string
			var ok bool
			var readObj interface{}
			obj := "object in database"
			t.Run("create operation", func(t *testing.T) {
				if id, ok = c.dbImp.Create(obj); !ok {
					t.Errorf("expected %t but got %t", true, ok)
				}
			})
			t.Run("retrieve operation", func(t *testing.T) {
				if readObj, ok = c.dbImp.Retrieve(id); !ok {
					t.Errorf("expected %t but got %t", true, ok)
				}
			})
		})
	}
}

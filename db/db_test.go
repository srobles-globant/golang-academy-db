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
		{name: "connection to FileDB", dbImp: &db.FileDb{FilePath: "./test.db"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("start not connected", func(t *testing.T) {
				if state := c.dbImp.Connected(); state {
					t.Errorf("expected connected state to be %t but got %t", false, state)
				}
			})
			t.Run("connect operation success", func(t *testing.T) {
				if ok := c.dbImp.Connect(); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
			})
			t.Run("connect operation failure", func(t *testing.T) {
				if ok := c.dbImp.Connect(); ok {
					t.Errorf("expected to return %t but got %t", false, ok)
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
			var id string
			var ok bool
			var readObj interface{}
			obj := "test object"
			obj2 := "another test object"
			t.Run("not connected operations failure", func(t *testing.T) {
				if _, ok = c.dbImp.Create(obj); ok {
					t.Errorf("expected Create to return %t but got %t", false, ok)
				}
				if _, ok = c.dbImp.Retrieve("dummy"); ok {
					t.Errorf("expected Retrieve to return %t but got %t", false, ok)
				}
				if ok = c.dbImp.Update("dummy", obj); ok {
					t.Errorf("expected Update to return %t but got %t", false, ok)
				}
				if ok = c.dbImp.Delete("dummy"); ok {
					t.Errorf("expected Delete to return %t but got %t", false, ok)
				}
			})
			t.Run("create operation success", func(t *testing.T) {
				c.dbImp.Connect()
				if id, ok = c.dbImp.Create(obj); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
			})
			t.Run("retrieve operation failure", func(t *testing.T) {
				if _, ok = c.dbImp.Retrieve("dummy"); ok {
					t.Errorf("expected to return %t but got %t", false, ok)
				}
			})
			t.Run("retrieve operation success", func(t *testing.T) {
				if readObj, ok = c.dbImp.Retrieve(id); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
				if convertedObj, convertedOk := readObj.(string); !convertedOk || convertedObj != obj {
					t.Errorf("expected to retrieve \"%s\" but got \"%v\"", obj, readObj)
				}
			})
			t.Run("update operation failure", func(t *testing.T) {
				if ok = c.dbImp.Update("dummy", obj2); ok {
					t.Errorf("expected to return %t but got %t", false, ok)
				}
			})
			t.Run("update operation success", func(t *testing.T) {
				if ok = c.dbImp.Update(id, obj2); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
				readObj, ok = c.dbImp.Retrieve(id)
				if convertedObj, convertedOk := readObj.(string); !convertedOk || convertedObj != obj2 {
					t.Errorf("expected to retrieve the updated object \"%s\" but got \"%v\"", obj2, readObj)
				}
			})
			t.Run("delete operation failure", func(t *testing.T) {
				if ok = c.dbImp.Delete("dummy"); ok {
					t.Errorf("expected to return %t but got %t", false, ok)
				}
			})
			t.Run("update operation success", func(t *testing.T) {
				if ok = c.dbImp.Delete(id); !ok {
					t.Errorf("expected to return %t but got %t", true, ok)
				}
				if _, ok = c.dbImp.Retrieve(id); ok {
					t.Errorf("expected retrieve to return %t but got %t", false, ok)
				}
			})
		})
	}
}

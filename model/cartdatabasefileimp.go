// Package model defines the structs of the app's data model and the interface to persist them
package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/srobles-globant/golang-academy-db/db"
)

// CartDatabaseFileImp implements service.CartDatabase with my custom Db interface
type CartDatabaseFileImp struct {
	db db.Db
}

// NewCartDatabaseFileImp creates a CartDatabaseImp with the provided Db
func NewCartDatabaseFileImp(db db.Db) *CartDatabaseFileImp {
	return &CartDatabaseFileImp{db: db}
}

// CreateCart create a cart object in the database
func (cd *CartDatabaseFileImp) CreateCart(c *Cart) (int, error) {
	if id, ok := cd.db.Create(*c); ok {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return intID, nil
		}
		return 0, fmt.Errorf("db: error writing new Cart. %w", err)
	}
	return 0, fmt.Errorf("db: error writing new Cart")
}

// GetCart retrieves a cart object from the database
func (cd *CartDatabaseFileImp) GetCart(id int) (*Cart, error) {
	if c, ok := cd.db.Retrieve(fmt.Sprintf("%d", id)); ok {
		cartv, ok2 := c.(Cart)
		if ok2 {
			return &cartv, nil
		}
		b, err := json.Marshal(c)
		if err != nil {
			return nil, fmt.Errorf("db: error retrieving the Cart %d. %w", id, err)
		}
		var cartp *Cart
		err = json.Unmarshal(b, cartp)
		if err != nil {
			return nil, fmt.Errorf("db: error retrieving the Cart %d. %w", id, err)
		}
		cartp.ID = id
		return cartp, nil
	}
	return nil, fmt.Errorf("db: error retrieving the Cart %d", id)
}

// UpdateCart updates an existing cart in the database with the provided cart
func (cd *CartDatabaseFileImp) UpdateCart(id int, c *Cart) error {
	if ok := cd.db.Update(fmt.Sprintf("%d", id), *c); ok {
		return nil
	}
	return fmt.Errorf("db: error updating the object")
}

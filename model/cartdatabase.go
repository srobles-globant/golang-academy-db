// Package model defines the structs of the app's data model and the interface to persist them
package model

import (
	"fmt"
	"strconv"

	"github.com/srobles-globant/golang-academy-db/db"
)

// CartDatabase defines the interface to persist carts
type CartDatabase interface {
	CreateCart(*Cart) (int, error)
	GetCart(int) (*Cart, error)
	UpdateCart(int, *Cart) error
}

// CartDatabaseImp implements CartDatabase with my custom Db interface
type CartDatabaseImp struct {
	db db.Db
}

// NewCartDatabaseImp creates a CartDatabaseImp with the provided Db an initiates it
func NewCartDatabaseImp(db db.Db) *CartDatabaseImp {
	return &CartDatabaseImp{db: db}
}

// CreateCart create a cart object in the database
func (cdi *CartDatabaseImp) CreateCart(c *Cart) (int, error) {
	if id, ok := cdi.db.Create(*c); ok {
		if intID, err := strconv.Atoi(id); err != nil {
			return intID, nil
		} else {
			return 0, fmt.Errorf("db: error writing new object. %w", err)
		}
	}
	return 0, fmt.Errorf("db: error writing new object")
}

// GetCart retrieves a cart object from the database
func (cdi *CartDatabaseImp) GetCart(id int) (*Cart, error) {
	return nil, nil
}

// UpdateCart updates an existing cart in the database with the provided cart
func (cdi *CartDatabaseImp) UpdateCart(id int, c *Cart) error {
	return nil
}

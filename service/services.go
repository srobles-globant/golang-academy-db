// Package service exposes the functions that handle the app's logic
package service

import "github.com/srobles-globant/golang-academy-db/model"

// CartDatabase defines the interface to persist carts
type CartDatabase interface {
	CreateCart(*model.Cart) (int, error)
	GetCart(int) (*model.Cart, error)
	UpdateCart(int, *model.Cart) error
}

// CartServiceImp implements router.CartService with a CartDatabase
type CartServiceImp struct {
	cd CartDatabase
}

// NewCartServiceImp creates a CartServiceImp with the provided CartDatabase
func NewCartServiceImp(cd CartDatabase) *CartServiceImp {
	return &CartServiceImp{cd: cd}
}

// CreateCart creates a new cart
func (cs *CartServiceImp) CreateCart(cart *model.Cart) error {
	return nil
}

// GetCart retrieves a cart
func (cs *CartServiceImp) GetCart(cartID int) (*model.Cart, error) {
	return nil, nil
}

// AddItems adds a slice of items to a cart
func (cs *CartServiceImp) AddItems(cartID int, items []model.Item) error {
	return nil
}

// ListItems retrieves a slice of all items of a cart
func (cs *CartServiceImp) ListItems(cartID int) ([]model.Item, error) {
	return nil, nil
}

// ChangeItemQuantity changes the quantity field of an item of a cart
func (cs *CartServiceImp) ChangeItemQuantity(cartID, itemID, quantity int) error {
	return nil
}

// RemoveItem removes the item from the cart
func (cs *CartServiceImp) RemoveItem(cartID, itemID int) error {
	return nil
}

// ClearCart removes all items from a cart
func (cs *CartServiceImp) ClearCart(cartID int) error {
	return nil
}

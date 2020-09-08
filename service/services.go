// Package service exposes the functions that handle the app's logic
package service

import (
	"fmt"

	"github.com/srobles-globant/golang-academy-db/model"
	"github.com/srobles-globant/golang-academy-db/util"
)

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
func (cs *CartServiceImp) CreateCart(cart *model.Cart) (int, error) {
	return cs.cd.CreateCart(cart)
}

// GetCart retrieves a cart
func (cs *CartServiceImp) GetCart(cartID int) (*model.Cart, error) {
	return cs.cd.GetCart(cartID)
}

// AddItems adds a slice of items to a cart
func (cs *CartServiceImp) AddItems(cartID int, items []model.Item) error {
	cart, err := cs.GetCart(cartID)
	if err != nil {
		return err
	}
	nextID := cart.Items[len(cart.Items)-1].ID + 1
	for _, item := range items {
		if itemSliceHasArticle(cart.Items, item.ArticleID) {
			continue
		}
		if _, err := util.GetArticle(item.ArticleID); err != nil {
			continue
		}
		item.ID = nextID
		nextID++
		cart.Items = append(cart.Items, item)
	}
	return cs.cd.UpdateCart(cartID, cart)
}

// ListItems retrieves a slice of all items of a cart
func (cs *CartServiceImp) ListItems(cartID int) ([]model.Item, error) {
	cart, err := cs.GetCart(cartID)
	if err != nil {
		return nil, err
	}
	return cart.Items, nil
}

// ChangeItemQuantity changes the quantity field of an item of a cart
func (cs *CartServiceImp) ChangeItemQuantity(cartID, itemID, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("Item quantity %d is invalid", quantity)
	}
	cart, err := cs.GetCart(cartID)
	if err != nil {
		return err
	}
	x := itemIndexFromID(cart.Items, itemID)
	if x == -1 {
		return fmt.Errorf("Item %d from cart %d doesn't exist", itemID, cartID)
	}
	cart.Items[x].Quantity = quantity
	return cs.cd.UpdateCart(cartID, cart)
}

// RemoveItem removes the item from the cart
func (cs *CartServiceImp) RemoveItem(cartID, itemID int) error {
	cart, err := cs.GetCart(cartID)
	if err != nil {
		return err
	}
	x := itemIndexFromID(cart.Items, itemID)
	if x == -1 {
		return fmt.Errorf("Item %d from cart %d doesn't exist", itemID, cartID)
	}
	cart.Items = append(cart.Items[:x], cart.Items[x+1:]...)
	return cs.cd.UpdateCart(cartID, cart)
}

// ClearCart removes all items from a cart
func (cs *CartServiceImp) ClearCart(cartID int) error {
	cart, err := cs.GetCart(cartID)
	if err != nil {
		return err
	}
	cart.Items = make([]model.Item, 0, 0)
	return cs.cd.UpdateCart(cartID, cart)
}

func itemSliceHasArticle(items []model.Item, articleID int) bool {
	for _, item := range items {
		if item.ArticleID == articleID {
			return true
		}
	}
	return false
}

func itemIndexFromID(items []model.Item, itemID int) int {
	x := -1
	for i, item := range items {
		if item.ID == itemID {
			x = i
			break
		}
	}
	return x
}

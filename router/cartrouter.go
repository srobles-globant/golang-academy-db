// Package router configures the app's router
package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/srobles-globant/golang-academy-db/model"
)

// CartService defines the methods to handle Carts
type CartService interface {
	CreateCart(cart *model.Cart) (int, error)
	GetCart(cartID int) (*model.Cart, error)
	AddItems(cartID int, items []model.Item) error
	ListItems(cartID int) ([]model.Item, error)
	ChangeItemQuantity(cartID, itemID, quantity int) error
	RemoveItem(cartID, itemID int) error
	ClearCart(cartID int) error
}

// CartRouter type holds the dependencies of the app's router
type CartRouter struct {
	cs CartService
}

// NewCartRouter creates a CartRouter with the provided CartService
func NewCartRouter(cs CartService) *CartRouter {
	return &CartRouter{cs: cs}
}

// Create a new cart
func (cr *CartRouter) createCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("=== create a cart ===")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, err, http.StatusInternalServerError, "Error processing the request")
		return
	}
	log.Println("body:\n" + string(b))
	var cart *model.Cart
	err = json.Unmarshal(b, cart)
	if err != nil {
		processError(w, err, http.StatusBadRequest, "Malformed JSON")
		return
	}
	id, err := cr.cs.CreateCart(cart)
	responseBody := model.ApiResponse{
		Message: "Cart created",
		Data:    struct{ id int }{id: id},
	}
	responseData, _ := json.Marshal(responseBody)
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

// Get a cart
func (cr *CartRouter) getCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("get a cart")
}

// Adding items to a cart
func (cr *CartRouter) addItems(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("add items to a cart")
}

// List all items of a specific cart
func (cr *CartRouter) listItems(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("list items of a cart")
}

// Changing the quantity of an existent item in a cart
func (cr *CartRouter) changeItemQuantity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("change the quantity of an item")
}

// Removing an item from a cart
func (cr *CartRouter) removeItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("remove item from a cart")
}

// Clear a specific cart (remove all items)
func (cr *CartRouter) clearCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("clear a cart")
}

// CreateRouter creates the app's router
func (cr *CartRouter) CreateRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/carts", cr.createCart)
	router.GET("/carts/:cart", cr.getCart)
	router.POST("/carts/:cart/items", cr.addItems)
	router.GET("/carts/:cart/items", cr.listItems)
	router.PATCH("/carts/:cart/items/:item", cr.changeItemQuantity)
	router.DELETE("/carts/:cart/items/:item", cr.removeItem)
	router.DELETE("/carts/:cart/items", cr.clearCart)

	return router
}

func processError(w http.ResponseWriter, err error, statusCode int, message string) {
	responseBody := model.ApiResponse{Message: message}
	log.Println(err)
	responseData, _ := json.Marshal(responseBody)
	w.WriteHeader(statusCode)
	w.Write(responseData)
}

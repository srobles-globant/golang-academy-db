package main

import (
	"log"
	"net/http"

	"github.com/srobles-globant/golang-academy-db/router"
)

func main() {
	r := &router.CartRouter{}
	log.Println("Running http handler in port 8080")
	log.Fatal(http.ListenAndServe(":8080", r.CreateRouter()))
}

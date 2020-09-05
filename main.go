package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/srobles-globant/golang-academy-db/db"
	"github.com/srobles-globant/golang-academy-db/model"
	"github.com/srobles-globant/golang-academy-db/router"
	"github.com/srobles-globant/golang-academy-db/service"
)

func main() {
	// Setup
	db := db.FileDb{FilePath: "./filedb.dat"}
	if ok := db.Connect(); !ok {
		log.Fatalln("Error conecting to the database")
	}
	cdb := model.NewCartDatabaseFileImp(&db)
	cs := service.NewCartServiceImp(cdb)
	cr := router.NewCartRouter(cs)

	// Start server
	server := &http.Server{Addr: ":8080", Handler: cr.CreateRouter()}
	go func() {
		log.Println("Running http server in port 8080")
		log.Println(server.ListenAndServe())
	}()

	// Prepare for listening the interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Stop the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println(server.Shutdown(ctx))

	// Close and cleanup
	db.Disconnect()
}

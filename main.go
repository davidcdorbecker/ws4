package main

import (
	"log"
	"net/http"
	"time"

	"ws4/controllers"

	"github.com/gorilla/mux"
)

func main() {

	//the router should be on a different package
	r := mux.NewRouter()
	r.HandleFunc("/api/character/{filter}/{toSearch}", controllers.Handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

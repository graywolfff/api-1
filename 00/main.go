package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{
		addr: ":8080",
	}
	// Initialize the ServeMux
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	log.Fatal(srv.ListenAndServe())

}

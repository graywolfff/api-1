package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) GetHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home Page"))
}
func (a *api) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("About Page"))
}

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

	mux.HandleFunc("GET /", api.GetHomePage)
	mux.HandleFunc("GET /about", api.GetAboutPage)

	log.Fatal(srv.ListenAndServe())

}

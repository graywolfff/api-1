package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Home page: /"))
			return
		case "/about":
			w.Write([]byte("About Page: /about"))
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
}

func main() {
	api := &api{
		addr: ":8080",
	}
	srv := &http.Server{Addr: api.addr, Handler: api}
	log.Fatal(srv.ListenAndServe())

}

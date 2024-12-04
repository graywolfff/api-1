package main

import "net/http"

func (app *application) healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

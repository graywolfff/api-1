package main

import (
	"encoding/json"
	"net/http"
)

type api struct {
	addr string
}

var users = []*User{}

func (a *api) getUsersHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// encode users slice to json
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

	// decode request body to User struct
	var payload *User
	decoder := json.NewDecoder(r.Body)
	// by default, Decode not raise error in case fields miss match.
	// if we need to make sure user payload is comfortable for data struct,
	// add `DisallowUnknownFields` to decoder.
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	u := &User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
}

package main

import (
	"net/http"
)

type HandlerF func(DB, http.ResponseWriter, *http.Request)

func SignIn(db DB, f HandlerF) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func makeHandlerfunc(db DB, w http.ResponseWriter, r *http.Request) {
	db.Store("Data")
}

func main() {
	s := &Store{}
	http.HandleFunc("/", SignIn(s, makeHandlerfunc))
	http.ListenAndServe(":8080", nil)
}

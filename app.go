package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// Initialize establishes a new router
func (a *App) Initialize(serviceName string) {
	a.Router = mux.NewRouter()
	a.initRoutes()
}

// Run starts serving the API
func (a *App) Run(addr string) {
	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal(err)
	}

}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/link", AddLink).Methods("POST")
	a.Router.HandleFunc("/link/{id}", GetLink).Methods("GET")

	a.Router.HandleFunc("/login", Login).Methods("POST")
	a.Router.HandleFunc("/logout", Logout).Methods("PUT")
}

func authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// GetLink gets a link from to a users collection
func GetLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// AddLink addes a new link to a users collection
func AddLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// RemoveLink removes a link from a users collection
func RemoveLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// UpdateLink changes the meta data on a users link
func UpdateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// ShareLink shares a link with another user
func ShareLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// Login creates a token and assigns it to a user
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// Logout invalidates a token
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// Register creates a mondb record for a new user
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

// use provides a cleaner interface for chaining middleware for single routes.
// Middleware functions are simple HTTP handlers (w http.ResponseWriter, r *http.Request)
//
//  r.HandleFunc("/login", use(loginHandler, rateLimit, csrf))
//  r.HandleFunc("/form", use(formHandler, csrf))
//  r.HandleFunc("/about", aboutHandler)
//
// See https://gist.github.com/elithrar/7600878#comment-955958 for how to extend it to suit simple http.Handler's
func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

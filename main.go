package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"goji.io"
	"goji.io/pat"
)

type status struct {
	Status string `json:"status"`
	Code string `json:"code"`
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/status"), handleStatus)
	mux.Use(logging)
	http.ListenAndServe("localhost:8989", mux)
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	currentStatus := status{
		Status: "OK",
		Code: "200",
	}

	jsonOut, _ := json.Marshal(currentStatus)
	fmt.Fprintf(w, string(jsonOut))
}

func logging(h http.Handler) http.Handler {  
    fn := func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Received request: %v\n", r.URL)
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}



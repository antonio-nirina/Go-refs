package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonio-nirina/go-refs/server/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware)
	router.Path("/").Methods("GET").HandlerFunc(controller.Handle(controller.CountriesHandler))
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
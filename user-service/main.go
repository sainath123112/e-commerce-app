package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath/e-commerce-app/user-service/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", handler.LoginHandler).Methods("POST")
	http.Handle("/", r)
	log.Println("Listening and serving on port: 8080")
	http.ListenAndServe(":8080", nil)
}

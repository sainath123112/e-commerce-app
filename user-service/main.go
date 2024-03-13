// User Service:
//
//	version: 0.1
//	title: User service Api
//
// Schemes: http, https
// Host:
// BasePath: /user/
//
//	Consumes:
//	- application/json
//
// Produces:
//   - application/json
//
// swagger:meta
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
	r.HandleFunc("/user/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/user/{userid}", handler.Authenticate(handler.GetUserDetails)).Methods("GET")
	http.Handle("/", r)
	log.Println("Listening and serving on port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

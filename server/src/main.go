package main

import (
	"github.com/gorilla/mux"
	"github.com/daleksprinter/catechdojo/controller"
	"github.com/daleksprinter/catechdojo/db"
	"net/http"
	"log"

)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", controller.CreateUserController).Methods("POST")
	r.HandleFunc("/user/get", controller.GetUserController).Methods("GET")

	return r
}

func main(){
	router := NewRouter()
	db.NewDB("root:password@tcp(catechdojo_db)/CATechDojo")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())

}


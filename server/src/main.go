package main

import (
	"github.com/daleksprinter/catechdojo/controller"
	"github.com/daleksprinter/catechdojo/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", controller.CreateUserController).Methods("POST")
	r.HandleFunc("/user/get", controller.GetUserController).Methods("GET")
	r.HandleFunc("/user/update", controller.UpdateUserController).Methods("PUT")

	return r
}

func main() {
	router := NewRouter()
	db.NewDB("root:password@tcp(catechdojo_db)/CATechDojo")
	defer db.DB.Close()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())

}

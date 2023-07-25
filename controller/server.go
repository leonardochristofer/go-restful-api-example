package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardochristofer/go-restful-api-example/controller/cakes"
	"github.com/leonardochristofer/go-restful-api-example/model"
)

var router *mux.Router

type cakeDatabase struct{}

func initHandlers() {
	server := &cakes.CakesServer{CakeDatabase: model.CakeRepository{}}

	router.HandleFunc("/api/cakes", server.GetAllCakes).Methods("GET")

	router.HandleFunc("/api/cakes/{id}", server.GetCake).Methods("GET")

	router.HandleFunc("/api/cakes", server.CreateCake).Methods("POST")

	router.HandleFunc("/api/cakes/{id}", server.UpdateCake).Methods("PUT")

	router.HandleFunc("/api/cakes/{id}", server.DeleteCake).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()
	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}

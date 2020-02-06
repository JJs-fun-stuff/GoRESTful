package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/catinello/base62"
)

func mainLogic(w http.ResponseWriter, r *http.Request){

}


func main(){
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	mainLogicHandler :=api.HandleFunc(mainLogic);
	api.Handle("/new").Methods(http.MethodPost)
	api.HandleFunc("/:url").Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorrilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request){
	// Fetch query parameters as a map
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id: %s", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s", queryParams["category"][0])
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")

	srv := &htttp.Server{
		Handler: r,
		Addr: "127.0.0.1.8000"
		WriteTimeout: 15 *time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
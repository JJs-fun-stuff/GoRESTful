package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/catinello/base62"
	"log"
	"encoding/json"
)

type Url struct {
	OriginalUrl string `json:"originalurl"`
	EncodedUrl string `json:"encodedurl"`
}

// GET /{url} Base62-encoded Url -> original Url
func returnURL(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	decodeUrl,err := base62.Decode(vars.url)
	if err != nil{
		log.Fatal("Cannot decode this given url.")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	addUrl := Url {
		OriginalUrl : decodeUrl,
		EncodedUrl: vars.url,
	}
	json.NewEncoder(w).Encode(addUrl)
	

}

// POST /new originalUrl -> base62-encoded Url
func createURL(w http.ResponseWriter, r *http.Request){
	
	// Decode a json data from HTTP body request
	url :=  Url{OriginalUrl : "", EncodedUrl : r.Body}
	url := json.NewDecoder(r.Body).Decode(url.OriginalUrl)
	encodeURL := base62.Encode(url)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Original URL is :%v", encodeURL)
}


func main(){
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/new",createURL).Methods("POST")
	r.HandleFunc("/{:url}", returnURL).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
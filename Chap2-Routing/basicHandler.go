package main

import (
	"log"
	"io"
	"net/http"
)

func Myserver(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello, world!\n")
}

func main(){
	http.HandleFunc("/hello", Myserver)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
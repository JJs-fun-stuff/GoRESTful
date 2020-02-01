package main

import(
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main(){
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("C:/Users/jojogelato/go/src/github.com/JJs-fun-stuff/Chap2-Routing/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
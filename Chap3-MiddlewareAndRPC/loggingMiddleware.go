package main

import(
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"os"
	"net/http"
)

func mainLogic(w http.ResponseWriter, r *http.Request){
	log.Println("Start processing request...")
	w.Write([]byte("OK"))
	log.Println("Finished processign request !!!")
	
}

func main(){
	r := mux.NewRouter()
	// make a router handle mainlogic
	r.HandleFunc("/", mainLogic)
	// make another router which handle logginghandler of mainlogic
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}
package main

import (
	"fmt"
	"net/http"
)

// accept http.Handler <- mainlogic as a argument 
func middleware(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase.")
		// Pass control back to the handler
		// Allow handler to execute handler logic
		handler.ServeHTTP(w,r)
		fmt.Println("Executing middleware AFTER request phase.")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request){
	//Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}


func main(){
	// HandlerFunc return a http handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/",middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}
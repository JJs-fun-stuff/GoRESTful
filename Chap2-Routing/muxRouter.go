package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n",vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

//Query-based matching
func QueryHandler(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()
	fmt.Fprintf(w, "Got parameter id : %s!",queryParams["id"])
	fmt.Fprintf(w, "Got parameter category : %s!\n", queryParams["category"])
}

func main(){
	r := mux.NewRouter()
	// r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler)
	
	//PathPrefix -> build a default path 
	// s := r.PathPrefix("/articles").SubRouter()
	// s.HandleFunc("{id}/settings", settingHandler)
	// s.HandleFunc("{id}/details", detailsHandler)
	
	//Strict Slash -> redirect ไปหา link แม่ 
	// r.StrictSlash(true)
	// r.Path("/articles/").Handler(ArticleHandler)
	
	// Encoded Path
	// r.UseEncodedPath()
	// r.NewRoute().Path("category/id") 

	//Revere mapping URL
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")
	// url, err := r.Get("articleRoute").URL("category","books","id","123")
	// fmt.Printf(url.URL)

	// Query-based matching
	// r.HandleFunc("/articles", QueryHandler)
	// r.Queries("id", "category")

	//Host-based matching
	//r.Host("aaa.bbb.ccc")
	//r.HandleFunc("/id1/id2/id3", MyHandler)
	// If we set like this -> http://aaa.bbb.ccc/111/222/333 will be matched 
	// We can match it with HTTP Methods(GET POST)


	//Enforce timeout for a server
	srv := &http.Server{
		Handler: r,
		Addr:"127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
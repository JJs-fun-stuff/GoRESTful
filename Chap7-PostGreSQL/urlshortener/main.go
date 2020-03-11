package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JJs-fun-stuff/Chap7-PostGreSQL/urlshortener/models"
	"github.com/catinello/base62"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type DBClient struct {
	db *sql.DB
}

// Model the Record stuff
type Record struct {
	ID  int `่่json:"id"`
	URL int `json:"url"`
}

// Method Get
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	id, _ := base62.Decode(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url where id=$1", id).Scan(&url)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

// create base62 link as url
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING ide, record.URL").Scan(&id)

	responseMap := map[string]interface{}{"encoded_string": base62.Encode(id)}

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create a new Router
	r := mux.NewRouter()
	//Attach a path with handler
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
		        		 		
package main

import(
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	session *mgo.Session
	collection *mgo.Collection
}

type Movie struct {
	ID 		bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	Name 	string 	`json:"name" bson:"name"`
	Year 	string 	`json:"year" bson:"year"`
	Directors 	[]string 	`json:"directors" bson:"directors"`
	Writers 	[]string 	`json:"writers" bson:"writers"`
	BoxOffice 	BoxOffice 	`json:"boxOffice" bson:"boxOffice"`
}
type BoxOffice struct {
	Budget 	uint64 	`json:"budget" bson:"budget"`
	Gross 	uint64 	`json:"gross" bson:"gross"`
}

// GET /v1/movies/{id:[a-z-A-Z0-9]*} fetches a movie to user with a given ID
func(db *DB) GetMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var movie Movie
	err := db.collection.Find(bson.M{"_id":bson.ObjectIdHex(vars["id"])}).One(&movie)

	if err !=nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type","application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

// POST /v1/movies : Add new movie to our MongoDB collection
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request){
	var movie Movie 
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	//Create Hash ID to insert
	movie.ID = bson.NewObjectId()
	err := db.Collection.insert(movie)
	
	if err !=nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type","application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

func main(){

	//------- Create a Database part -------//
	session, err := mgo.Dial("127.0.0.1")
	//Fetch a collection using DB and C chaining together
	c := session.DB("appdb").C("movies")
	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//------- Create a Router part -------//
	r := mux.NewRouter()
	r.HandleFunc("/v1/movies/{id:[a-z-A-Z0-9]*}",db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies",db.PostMovie).Methods("POST")

	srv := &http.Server {
		Handler: r,
		Addr:"127.0.0.1:8000",
		//Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
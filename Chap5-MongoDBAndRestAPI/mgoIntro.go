package main

import(
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Name string `bson:"name"`
	Year string `bson:"year"`
	Directors string `bson:"directors"`
	Writers string `bson:"writers"`
	BoxOffice `bson:"boxOffice`

}

type boxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`

}


func main() {
	//Create session
	session, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}
	defer session.Close()
	//Fetch a collection using DB and C chaining together
	c := session.DB("appdb").C("movies")


	darkNight := &Movie{

	}
	// Insert  a data to DB similar to POST method
	err = c.Insert(darkNight)

	if err != nil{
		log.Fatal(err)
	}

	result := Movie{}
	// read  a data from DB similar to GET method
	//bson.M use to translate go query to code that mongoDB understands
	//
	err := c.Find(bson.M{"boxOffice.budget": bson.M{"$gt":150000000}}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", result.Name)
}

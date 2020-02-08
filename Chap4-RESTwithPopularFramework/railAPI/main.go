package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"C:\Users\jojogelato\go\src\github.com\JJs-fun-stuff\Chap4-RESTwithPopularFramework\dbutils"

	"encoding/json"
	"net/http"
	"time"
	"github.com/emicklei/go-restful"
)

func main(){
	var DB *sql.DB
///////////////////// Resource struct //////////////////
	type TrainResource struct {
		ID int
		DriverName string
		OperatingStatus bool
	}

	type StationnResource struct {
		ID int
		Name string
		OpeningTime time.Time
		ClosingTime time.Time
	}

	type ScheduleResource struct {
		ID int 
		TrainID int
		StationID int
		ArrivalTime time.Time
	}

	// Connect to the database
	DB, err := sql.Open("sqlite3","./railapi.db")

	if err != nil{
		log.Println("Driver creation failed!")
	}

	// Create tables
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	ws.Conatainer.Router(restful.CurlyRouter{})

	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")
	server := &http.Server{ Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func (t *TrainResource) Register(container *restful.Container){

	// Initialize a web service
	ws := new(restful.WebService)
	ws.Path("/v1/trains")
	Consumes(restful.MIME_JSON)
	Produces(restful.MIME_JSON)


	// Routing many path to different handler
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.deleteTrain))
	

	// Add the service to application
	container.Add(ws)
}

// GET localhost:8000/v1/trains/1
func(t *TrainResource) getTrain(request *restful.Request, response *restful.Response){
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?",id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	
	if err != nil{
		response.AddHeader("Content-Type","text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response){
	log.Println(request.Request.Body)

	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	// Error handling is obvious here. So omitting...
	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values(?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)

	if err == nil{
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else{
		response.AddHeader("Content-Type","text/plain")
		response.WriteErrorString(http.StatuseInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) deleteTrain(request *restful.Request, response *restful.Response){
	id := request.PathParameter("train-id")
	statement, err = DB.Prepare("DELETE from train WHERE id =?")
	statement.Exec(id)
	if err == nil{
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type","text/plain")
		response.WriteErrorString(http.StatuseInternalServerError, err.Error())
	}
}


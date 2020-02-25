package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"C:\\Users\\jojogelato\\go\\src\\github.com\\JJs-fun-stuff\\Chap4-RESTwithPopularFramework\\dbutils"
)

var DB *sql.DB

type StationResource struct{
	ID int `json:"id"`
	Name string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// Get the station details
// GET localhost:8000/v1/trains/1

func GetStation (c *gin.Context){
	var station StationResource
	id :=  c.Param("station_id")
	err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result" : station,
		})
	}
}

// POST http://localhost:8000/v1/trains
func CreateStation (c *gin.Context){
	var station StationResource
	// Parse the body into our resource
	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?,?,?)")

		result, _ := statement.</span>Exec(station.Name, station.OpeningTime, station.ClosingTime)
		
		if err ==nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOk, gin.H{
				"result": station,
			})
			} else {
				c.String(http.StatusInternalServerError, err.Error())
			}
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	}

// DELETE http://localhost:8000/v1/trains/1
func RemoveStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station-id")
	statement,_ := DB.Prepare("DELETE from station WHERE id =?")
	_,err := statement.Exec(id)

	if err != nil{
		c.JSON(500,gin.H{
			"error":err.Error())
	} else {
		c.String(http.StatusOK,"")

	}
}
func main(){
	var err error
	DB, err = sql.Open("sqlite3","./railapi.db")

	if err != nil{
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)

	r := gin.Default()
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8000") // Default listen and serve on 0.0.0.0:8000

}
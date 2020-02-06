package main
import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)
type city struct {
	Name string
	Area uint64
}


func mainLogic(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		var tempCity city
		// decode the body of the request
		decoder := json.NewDecoder(r.Body)
		// assigned the decoded value to tempCity
		err := decoder.Decode(&tempCity)
		if err != nil{
			panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main(){
	http.HandleFunc("/city", mainLogic)
	http.ListenAndServe(":8000", nil)
}
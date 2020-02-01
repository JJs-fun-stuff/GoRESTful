package main
import (
	"log"
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string)string {
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil{
		log.Fatal(fmt.Sprint(err)+": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil{
		log.Fatal(fmt.Sprint(err)+": " + stderr.String())
	}
	return out.String()
}
// Similar to http.HandleFunc but have additional argument -> params httprouter.Params
func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	fmt.Fprintf(w, getCommandOutput("C:/Users/jojogelato/go/bin","version"))
}


func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	fmt.Fprintf(w, getCommandOutput("C:/Users/jojogelato/go/src/github.com/JJs-fun-stuff/Chap2-Routing", params.ByName("name")))
}

func main(){
	//create new router
	//Able to handle GET POST DELETE
	router := httprouter.New()
	//GET method able to takes two argument "url" and handler func
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}


package main
import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer)GiveServerTime(args *Args, reply *int64) error{
	//Fill the pointer to send the data back
	*reply = time.Now().Unix()
	return nil

}

func main(){
	timeserver := new(TimeServer)
	// the rpc server intend to export object type Timserver(int64)
	rpc.Register(timeserver)

	//HandleHTTP registers HTTP handler for RPC messages to "DefaultServer"
	rpc.HandleHTTP()
	l, e := net.Listen("tcp",":1234")

	if e != nil{
		log.Fatal("listen error:",e)
	}

	//http.Serve use to serve it as a running program
	http.Serve(l,nil)
}

curl -X POST http://localhost:8000/rpc -H 'cache-control: no-cache' -H 'content-type: application/json'  -d '{"method": "JSONServer.GiveBookDetail", "params": [{"Id": "1234"}],"id": "1"}'
package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"

	"github.com/tcp-x/cd-rpc/service"
)

var server = flag.String("server_port", "localhost:9999", "Address at which to reach the server.")

// var factor1 = flag.Int("factor1", 3, "First factor to multiply.")
// var factor2 = flag.Int("factor2", 4, "Second factor to multiply.")
var req = flag.String("req", `{"ctx":"sys", "m":"User", "c": "User", "dat":[]}`, "request.")

func main() {
	flag.Parse()

	client, err := rpc.DialHTTP("tcp", *server)
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	// request := &service.MultiplicationRequest{*factor1, *factor2}
	// response := new(service.MultiplicationResponse)

	request := &service.CdRequest{*req}
	response := new(service.CdResponse)

	// err = client.Call("MultiplicationService.Multiply", request, &response)
	// if err != nil {
	// 	log.Fatal("MultiplicationService error:", err)
	// }
	err = client.Call("UserService.Auth", request, &response)
	if err != nil {
		log.Fatal("MultiplicationService error:", err)
	}

	fmt.Println(response.Resp)
}

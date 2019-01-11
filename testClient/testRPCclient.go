package main

import (
	"net"
	"net/rpc/jsonrpc"
)

// структура
type Args struct {
	NAME string
	UID  string
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	conn := jsonrpc.NewClient(client)
	defer conn.Close()

	args := &Args{"Ivan", "-"}

	var reply string

	err = conn.Call("RPCServer.Add", args, &reply)
	checkError(err)
	//fmt.Printf("Result: ", reply)
	args = &Args{"change", reply}
	err = conn.Call("RPCServer.Update", args, &reply)
	checkError(err)
	//fmt.Printf("Result: ", reply)
	err = conn.Call("RPCServer.Show", args, &reply)
	checkError(err)
	//fmt.Printf("Result: ", reply)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

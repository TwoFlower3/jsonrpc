package main

import (
	"database/sql"
	"net"
	"net/rpc"
	"testing"
)

type TestRPCServer struct{}

func TestServer(t *testing.T) {

	_, err := sql.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")

	if err != nil {
		t.Fatal("Error connection to db")
	}

	call := new(TestRPCServer)

	server := rpc.NewServer()
	server.Register(call)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal("Error create server")
	}

	defer listen.Close()
}

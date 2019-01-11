package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	_ "github.com/lib/pq"
)

// инициализация
var db *sql.DB

// структура
type Args struct {
	NAME string
	UID  string
}

// структура
type RPCServer struct{}

// Функция добавления записи
func (t *RPCServer) Add(args *Args, res *string) error {

	var returnid string

	fmt.Println("args: ", args.NAME)
	err := db.QueryRow("INSERT INTO public.user (login) VALUES ($1) returning id;", args.NAME).Scan(&returnid)
	*res = returnid
	if err != nil {
		panic(err)
	}

	return nil
}

// Функция обновления записи
func (t *RPCServer) Update(args *Args, res *string) error {

	fmt.Println("args: ", args.NAME, args.UID)
	_, err := db.Exec("update public.user set login=$1 where id=$2", args.NAME, args.UID)

	if err != nil {
		panic(err)
	}

	*res = "success"
	return nil
}

// Функция вывода
func (t *RPCServer) Show(args *Args, res *string) error {

	result, err := db.Query("SELECT * FROM public.user")

	if err != nil {
		panic(err)
	}

	fmt.Println("result: ", result)

	for result.Next() {
		var name string
		var date time.Time
		var id string
		err = result.Scan(&name, &id, &date)
		if err != nil {
			panic(err)
		}
		fmt.Println("show: ", name, date, id)
	}

	*res = "success"
	return nil
}

func main() {

	var err error

	// инициализация базы данных
	db, err = sql.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	call := new(RPCServer)

	server := rpc.NewServer()
	server.Register(call)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	config, router := server.InitServer()
	isServerOnline := server.StartServer(config.Port, router)
	fmt.Println("Server status =" + isServerOnline + "")

}

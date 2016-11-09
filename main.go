package main

import (
	"flag"
	"fmt"

	"web-app/server"
)

var run = flag.Bool("run", true, "Flag to run server")

func main() {
	flag.Parse()
	config, router := server.initServer(*run)
	isServerOnline := server.startServer(config.Port, router)
	fmt.Println("Server status =" + isServerOnline + "")

}

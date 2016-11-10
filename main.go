package main

import (
	"log"

	"github.com/web-app-go/config"
	"github.com/web-app-go/server"
	"github.com/web-app-go/stringutil"
)

func main() {
	// String
	log.Printf(stringutil.Reverse("!oG ,olleH"))

	// Read from config
	conf := config.ReadConfigFile()

	// Start server
	server.InitServer()
	server.StartServer(conf.Port)

}

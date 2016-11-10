package main

import (
	"fmt"

	"github.com/web-app-go/config"
	"github.com/web-app-go/logger"
	"github.com/web-app-go/server"
	"github.com/web-app-go/stringutil"
)

func main() {
	// logger
	logger.SetLogger()

	// String
	logger.GetNotice(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("")

	// Read from config
	conf := config.ReadConfigFile()

	// Start server
	server.InitServer()
	server.StartServer(conf.Port)

}

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/web-app-go/logger"
)

// Config object that holds the information stored in config.json
type Config struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Custom   Custom `json:"custom"`
	Port     string `json:"port"`
}

type Custom struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

func ReadConfigFile() (config Config) {
	var conf Config

	// Read from config
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Error loading config:", err)
		return
	}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Println("Error unmarshalling config:", err)
		return
	}

	toString(conf)
	return conf

}

// Print
func toString(config Config) {
	logger.GetNotice("Success in getting the config file in a better format:")
	logger.GetDebugf("\t=> Name:", config.Name)
	logger.GetDebugf("\t=> Email:", config.Email)
	logger.GetDebugf("\t=> Password:", config.Password)
	logger.GetDebugf("\t=> Custom.First:", config.Custom.First)
	logger.GetDebugf("\t=> Custom.Second:", config.Custom.Second)
	logger.GetDebugf("\t=> Port:", config.Port)
	fmt.Println("")
}

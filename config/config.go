package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	log.Println("Success in getting the config file", conf)
	toString(conf)
	return conf

}

// Print
func toString(config Config) {
	log.Println("Success in getting the config file in a better format:")
	fmt.Println("\t=> Name:", config.Name)
	fmt.Println("\t=> Email:", config.Email)
	fmt.Println("\t=> Password:", config.Password)
	fmt.Println("\t=> Custom.First:", config.Custom.First)
	fmt.Println("\t=> Custom.Second:", config.Custom.Second)
	fmt.Println("\t=> Port:", config.Port)
	fmt.Println("")
}

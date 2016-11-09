package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// LoadConfig reads the config.json and loads the information in the config object
func loadConfig() (conf Config, err error) {
	// Read the config file.
	read, err := ioutil.ReadFile("conf.json")
	if err != nil {
		fmt.Println("Err: can't read the config file")
		panic(err)
	}
	// Unmarshal the JSON
	err = json.Unmarshal(read, &conf)
	if err != nil {
		fmt.Println("Err: parse the json to a struct")
		panic(err)
	}

	fmt.Println("Ok: values retrieved:", config)
	return conf, nil
}

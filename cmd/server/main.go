package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/AlexBaykov/go-rest-api/internal/app/server"
)

var (
	configPath string
)

const (
	defaultConfigPath = "configs/config.json"
)

func init() {
	flag.StringVar(&configPath, "config-path", defaultConfigPath, "Path to config file")
}

func main() {
	flag.Parse()

	byteConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error while reading config file!")
	}

	config := &server.Config{}
	err = json.Unmarshal(byteConfig, config)
	if err != nil {
		log.Fatal("Error while reading config file!")
	}

	srv := server.NewServer()
	if err := srv.Start(config); err != nil {
		log.Fatal("Error while starting the server!")
	}

}

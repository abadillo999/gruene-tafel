package main

import (
	"fmt"
	"flag"
	"config"
	"server"
	"processor"
)

func main() {
	configPathPtr := flag.String("config", "", "Path to the json config file.")
    flag.Parse()

    fmt.Println("Configuration file path:", *configPathPtr)

	config := config.NewConfig(*configPathPtr)
	config.GetDBConfig()

	processor := processor.NewProcessor(config.GetENVConfig())

	server := server.NewServer(config.GetServerConfig(), processor)
	server.Run()
}
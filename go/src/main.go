package main

import (
	//"config"
	"server"
)

func main() {

	config := config.neWConfig("/etc/grueneConfig.json")

	processor := Processor.newProcessor(config)
	server := &server.server{}
	server.initialize(config)
	server.run(":3000")
}
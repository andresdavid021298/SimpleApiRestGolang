package main

import (
	"ApiRest/configuration"
	"log"
)

func main() {
	configuration.RestConfig()
	log.Println("Start App")
}

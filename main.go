package main

import (
	"fmt"
	"services.com/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8995")
}

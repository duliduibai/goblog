package main

import (
	"fmt"
	"github.com/duliduibai/goblog/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8995")
}

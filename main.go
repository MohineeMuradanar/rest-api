package main

import (
	"AssDeploy/controller"
	"fmt"
)

//go:generate swagger generate spec -m -o ./swagger.json
func main() {

	fmt.Println("Starting the application....")

	controller.Handler()

}

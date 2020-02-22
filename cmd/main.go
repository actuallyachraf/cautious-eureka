package main

import (
	"github.com/actuallyachraf/backend-challenge/trendy/api"
)

// Main function only launches the main go routines for serving the webservice
// and listening for system interrupts such as Ctrl + C.
func main() {
	service := api.NewService()
	service.SetupService()
	service.Start()
}

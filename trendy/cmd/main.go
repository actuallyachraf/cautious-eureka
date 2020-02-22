package main

import (
	"github.com/actuallyachraf/backend-challenge/trendy/api"
)

func main() {
	service := api.NewService()
	service.SetupService()
	service.Start()
}

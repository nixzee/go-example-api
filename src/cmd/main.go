package main

import (
	"log"

	bootstrap "github.com/kardianos/service"
	service "github.com/nixzee/go-example-api/internal/service"
)

var (
	Version     string = "dev"
	Commit      string = "00000000"
	Name        string = "go-example-api"
	DisplayName string = "Golang Example API"
	Description string = ""
)

func main() {
	// Service configuration
	config := &bootstrap.Config{
		Name:        Name,
		DisplayName: DisplayName,
		Description: Description,
	}

	// Get the implemented service (program)
	program := service.NewProgam(Version, Commit)

	// Create the instance
	instance, err := bootstrap.New(program, config)
	if err != nil {
		log.Fatal(err)
	}

	// Run the instance
	err = instance.Run()
	if err != nil {
		log.Fatal(err)
	}
}

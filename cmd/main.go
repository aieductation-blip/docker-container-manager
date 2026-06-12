package main

import (
	"flag"
	"fmt"

	"docker-container-manager/internal/handlers"
)

func main() {
	// Define command-line flags
	listPtr := flag.Bool("list", false, "List all Docker containers")
	startPtr := flag.String("start", "", "Start a Docker container by name or ID")
	stopPtr := flag.String("stop", "", "Stop a Docker container by name or ID")

	flag.Parse()

	// Initialize handler
	handler := handlers.NewHandler()

	// Process flags
	if *listPtr {
		handler.ListContainers()
	} else if *startPtr != "" {
		handler.StartContainer(*startPtr)
	} else if *stopPtr != "" {
		handler.StopContainer(*stopPtr)
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
}

package handlers

import (
	"fmt"

	"docker-container-manager/internal/models"
	"docker-container-manager/internal/docker"
)

// Handler struct encapsulates methods for container management

type Handler struct {
	dockerClient *docker.Client
}

// NewHandler initializes and returns a new Handler
func NewHandler() *Handler {
	client, err := docker.NewClient()
	if err != nil {
		fmt.Println("Error initializing Docker client:", err)
	}
	return &Handler{dockerClient: client}
}

// ListContainers lists all Docker containers
func (h *Handler) ListContainers() {
	containers, err := h.dockerClient.ListContainers()
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return
	}

	fmt.Println("Running Docker Containers:")
	for _, container := range containers {
		fmt.Printf("- ID: %s, Names: %v, Status: %s\n", container.ID, container.Names, container.Status)
	}
}

// StartContainer starts a container by name or ID
func (h *Handler) StartContainer(id string) {
	err := h.dockerClient.StartContainer(id)
	if err != nil {
		fmt.Printf("Error starting container %s: %v\n", id, err)
	} else {
		fmt.Printf("Successfully started container %s\n", id)
	}
}

// StopContainer stops a container by name or ID
func (h *Handler) StopContainer(id string) {
	err := h.dockerClient.StopContainer(id)
	if err != nil {
		fmt.Printf("Error stopping container %s: %v\n", id, err)
	} else {
		fmt.Printf("Successfully stopped container %s\n", id)
	}
}

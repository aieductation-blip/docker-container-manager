package models

// Container represents a Docker container

type Container struct {
	ID     string
	Names  []string
	Image  string
	Status string
}

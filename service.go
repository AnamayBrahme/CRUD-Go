package main

import "fmt"

// Service represents a cloud service.
type Service struct {
	ID          int
	Name        string
	Environment string
	Status      string
}

// CREATE
func createService(services map[int]Service, service Service) {
	services[service.ID] = service
	fmt.Println("Service created successfully!")
}

// READ
func listServices(services map[int]Service) {
	for _, service := range services {
		fmt.Println("--------------------")
		fmt.Println("ID:", service.ID)
		fmt.Println("Name:", service.Name)
		fmt.Println("Environment:", service.Environment)
		fmt.Println("Status:", service.Status)
	}
}

// UPDATE
func updateService(services map[int]Service, id int, newStatus string) {
	service, exists := services[id]

	if !exists {
		fmt.Println("Service not found!")
		return
	}

	service.Status = newStatus
	services[id] = service

	fmt.Println("Service updated successfully!")
}

// DELETE
func deleteService(services map[int]Service, id int) {
	_, exists := services[id]

	if !exists {
		fmt.Println("Service not found!")
		return
	}

	delete(services, id)
	fmt.Println("Service deleted successfully!")
}

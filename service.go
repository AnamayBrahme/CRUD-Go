package main

import "fmt"

type Service struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

// CREATE
func createService(services map[int]Service, service Service) {
	services[service.ID] = service
	fmt.Println("Service created successfully!")
}

// READ
func listServices(services map[int]Service) []Service {
	var result []Service

	for _, service := range services {
		result = append(result, service)
	}

	return result
}

// UPDATE
func updateService(services map[int]Service, id int, newService Service) bool {
	_, exists := services[id]

	if !exists {
		return false
	}

	newService.ID = id
	services[id] = newService

	return true
}

// DELETE
func deleteService(services map[int]Service, id int) bool {
	_, exists := services[id]

	if !exists {
		return false
	}

	delete(services, id)
	return true
}

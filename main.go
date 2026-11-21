package main

import "fmt"

func main() {

	// Create temporary storage.
	services := make(map[int]Service)

	// CREATE
	createService(services, Service{
		ID:          1,
		Name:        "Cloud Monitor",
		Environment: "Development",
		Status:      "Running",
	})

	createService(services, Service{
		ID:          2,
		Name:        "Database API",
		Environment: "Production",
		Status:      "Stopped",
	})

	// READ
	fmt.Println("\nAll Services:")
	listServices(services)

	// UPDATE
	fmt.Println("\nUpdating Service 1:")
	updateService(services, 1, "Stopped")

	// DELETE
	fmt.Println("\nDeleting Service 2:")
	deleteService(services, 2)

	// READ AGAIN
	fmt.Println("\nRemaining Services:")
	listServices(services)
}

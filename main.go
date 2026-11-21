package main

import (
	"fmt"
	"net/http"
)

var services = make(map[int]Service)

func main() {

	// Add example services.
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "index.html")
	})

	// Register CRUD routes.
	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getServicesHandler(w, r)

		case http.MethodPost:
			createServiceHandler(w, r)

		case http.MethodPut:
			updateServiceHandler(w, r)

		case http.MethodDelete:
			deleteServiceHandler(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8083", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

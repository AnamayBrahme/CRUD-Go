package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GET: Retrieve all services.
func getServicesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	servicesList := listServices(services)

	json.NewEncoder(w).Encode(servicesList)
}

// POST: Create a new service.
func createServiceHandler(w http.ResponseWriter, r *http.Request) {

	var service Service

	err := json.NewDecoder(r.Body).Decode(&service)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if service.ID == 0 {
		http.Error(w, "Service ID is required", http.StatusBadRequest)
		return
	}

	if _, exists := services[service.ID]; exists {
		http.Error(w, "Service ID already exists", http.StatusConflict)
		return
	}

	createService(services, service)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(service)
}

// PUT: Update an existing service.
func updateServiceHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	var service Service

	err = json.NewDecoder(r.Body).Decode(&service)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updated := updateService(services, id, service)

	if !updated {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services[id])
}

// DELETE: Remove a service.
func deleteServiceHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	deleted := deleteService(services, id)

	if !deleted {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

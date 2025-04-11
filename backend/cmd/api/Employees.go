package main

import (
	"backend/internal/store"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	employees, err := app.Store.Employee.GetAllEmployees(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (app *application) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Firstname  string `json:"firstname"`
		Lastname   string `json:"lastname"`
		Role       string `json:"role"`
		Department string `json:"department"`
		Email      string `json:"email"`
		ContactNo  string `json:"contact_no"`
		Manager    string `json:"manager"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee := store.Employee{
		Firstname:  payload.Firstname,
		Lastname:   payload.Lastname,
		Role:       payload.Role,
		Department: payload.Department,
		Email:      payload.Email,
		ContactNo:  payload.ContactNo,
		Manager:    payload.Manager,
	}

	err := app.Store.Employee.Create(r.Context(), employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee created"})
}

func (app *application) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Firstname  string `json:"firstname"`
		Lastname   string `json:"lastname"`
		Role       string `json:"role"`
		Department string `json:"department"`
		Email      string `json:"email"`
		ContactNo  string `json:"contact_no"`
		Manager    string `json:"manager"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedEmployee := store.Employee{
		Firstname:  payload.Firstname,
		Lastname:   payload.Lastname,
		Role:       payload.Role,
		Department: payload.Department,
		Email:      payload.Email,
		ContactNo:  payload.ContactNo,
		Manager:    payload.Manager,
	}

	err = app.Store.Employee.UpdateEmployee(r.Context(), objectID, updatedEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated"})
}

func (app *application) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = app.Store.Employee.DeleteEmployee(r.Context(), objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted"})
}

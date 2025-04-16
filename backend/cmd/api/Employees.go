package main

import (
	"backend/internal/store"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type EmployeePayload struct {
	Firstname  *string `json:"firstname"`
	Lastname   *string `json:"lastname"`
	Role       *string `json:"role"`
	Department *string `json:"department"`
	Email      *string `json:"email"`
	ContactNo  *string `json:"contact_no"`
	Manager    *string `json:"manager"`
}

func validatePayload(payload *EmployeePayload) error {
	if payload.Firstname == nil || *payload.Firstname == "" ||
		payload.Lastname == nil || *payload.Lastname == "" ||
		payload.Email == nil || *payload.Email == "" {
		return errors.New("firstname, lastname, and email are required")
	}
	return nil
}

func (app *application) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	pq := store.ParsePagination(r)

	employees, err := app.Store.Employee.GetAllEmployees(r.Context(), pq)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	app.jsonResponse(w, http.StatusOK, employees)
}

func (app *application) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var payload EmployeePayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateEmployeePayload(&payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := app.Store.Employee.IsEmailExists(r.Context(), *payload.Email)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if exists {
		writeJSONError(w, http.StatusConflict, "Email already exists")
		return
	}

	emp := store.Employee{
		Firstname:  *payload.Firstname,
		Lastname:   *payload.Lastname,
		Role:       valueOrEmpty(payload.Role),
		Department: valueOrEmpty(payload.Department),
		Email:      *payload.Email,
		ContactNo:  valueOrEmpty(payload.ContactNo),
		Manager:    valueOrEmpty(payload.Manager),
	}

	if err := app.Store.Employee.Create(r.Context(), emp); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to create employee")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"message": "Employee created"})
}

func (app *application) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing id parameter")
		return
	}

	objectID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var payload EmployeePayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	update := buildUpdateMap(&payload)

	if len(update) == 0 {
		writeJSONError(w, http.StatusBadRequest, "No fields to update")
		return
	}

	if err := app.Store.Employee.UpdateEmployee(r.Context(), objectID, update); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to update employee")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Employee updated"})
}

func (app *application) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing id parameter")
		return
	}

	objectID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = app.Store.Employee.DeleteEmployee(r.Context(), objectID)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Employee deleted"})
}

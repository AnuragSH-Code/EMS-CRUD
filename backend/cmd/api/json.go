package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func validateEmployeePayload(payload *EmployeePayload) error {
	if payload.Firstname == nil || *payload.Firstname == "" ||
		payload.Lastname == nil || *payload.Lastname == "" ||
		payload.Email == nil || *payload.Email == "" {
		return errors.New("firstname, lastname, and email are required")
	}
	return nil
}

func valueOrEmpty(val *string) string {
	if val != nil {
		return *val
	}
	return ""
}

func buildUpdateMap(payload *EmployeePayload) bson.M {
	update := bson.M{}
	if payload.Firstname != nil {
		update["firstname"] = *payload.Firstname
	}
	if payload.Lastname != nil {
		update["lastname"] = *payload.Lastname
	}
	if payload.Email != nil {
		update["email"] = *payload.Email
	}
	if payload.Role != nil {
		update["role"] = *payload.Role
	}
	if payload.Department != nil {
		update["department"] = *payload.Department
	}
	if payload.ContactNo != nil {
		update["contact_no"] = *payload.ContactNo
	}
	if payload.Manager != nil {
		update["manager"] = *payload.Manager
	}
	return update
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {

	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelope{Error: message})
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return writeJSON(w, status, &envelope{Data: data})
}

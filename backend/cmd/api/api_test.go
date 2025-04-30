package main

import (
	"backend/internal/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestApp(mockStore *store.MockEmployeeStore) *application {
	return &application{
		Store:  &store.Storage{Employee: mockStore},
		Config: config{addr: ":8080"},
	}
}

func strPtr(s string) *string { return &s }

func TestGetAllEmployees_Success(t *testing.T) {
	mockStore := &store.MockEmployeeStore{
		Employees: []*store.Employee{
			{Firstname: "Harshit", Lastname: "Saxena", Email: "harshit@saxena.in"},
		},
	}
	app := setupTestApp(mockStore)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var response struct {
		Data []store.Employee `json:"data"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if len(response.Data) != 1 {
		t.Errorf("expected 1 employee, got %d", len(response.Data))
	}
}

func TestGetAllEmployees_DatabaseError(t *testing.T) {
	mockStore := &store.MockEmployeeStore{ShouldError: true}
	app := setupTestApp(mockStore)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestCreateEmployee_Success(t *testing.T) {
	mockStore := &store.MockEmployeeStore{}
	app := setupTestApp(mockStore)

	payload := map[string]interface{}{
		"firstname":  "Nausheen",
		"lastname":   "Javed",
		"role":       "People Team",
		"department": "HR",
		"email":      "nausheen@javed.in",
		"contact_no": "1234567890",
		"manager":    "Stuti Tyagi",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/v1/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
}

func TestCreateEmployee_ValidationError(t *testing.T) {
	app := setupTestApp(&store.MockEmployeeStore{})

	payload := map[string]interface{}{"firstname": "Vaibhav"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/v1/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}

func TestUpdateEmployee_Success(t *testing.T) {
	mockStore := &store.MockEmployeeStore{}
	app := setupTestApp(mockStore)

	payload := map[string]interface{}{
		"firstname": "Sachin",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("PUT", "/v1/employees?id=507f1f77bcf86cd799439011", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestUpdateEmployee_InvalidID(t *testing.T) {
	app := setupTestApp(&store.MockEmployeeStore{})

	req := httptest.NewRequest("PUT", "/v1/employees?id=invalid", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}

func TestDeleteEmployee_Success(t *testing.T) {
	mockStore := &store.MockEmployeeStore{}
	app := setupTestApp(mockStore)

	req := httptest.NewRequest("DELETE", "/v1/employees?id=507f1f77bcf86cd799439011", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestDeleteEmployee_DatabaseError(t *testing.T) {
	mockStore := &store.MockEmployeeStore{DeleteError: true}
	app := setupTestApp(mockStore)

	req := httptest.NewRequest("DELETE", "/v1/employees?id=507f1f77bcf86cd799439011", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestCorsMiddleware(t *testing.T) {
	app := setupTestApp(&store.MockEmployeeStore{})
	req := httptest.NewRequest("OPTIONS", "/v1/employees", nil)
	rr := httptest.NewRecorder()
	app.mount().ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("CORS headers not set properly")
	}
}

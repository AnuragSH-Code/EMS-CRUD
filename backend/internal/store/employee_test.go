// File: internal/store/employee_test.go
package store

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var testStore *MongoEmployeeStore

func setupTestDB(t *testing.T) {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	testStore = &MongoEmployeeStore{
		collection: client.Database("ems-brevo-test").Collection("stud"),
	}
}

func TestCreateAndGetEmployee(t *testing.T) {
	if os.Getenv("SKIP_DB_TESTS") == "true" {
		t.Skip("Skipping DB tests in CI")
	}

	setupTestDB(t)
	testStore.collection.Drop(context.TODO())

	tempEmp := Employee{
		Firstname:  "Nausheen",
		Lastname:   "Javed",
		Email:      "Nausheen@Javed.com",
		Department: "HR",
		Role:       "People",
		ContactNo:  "1234567890",
		Manager:    "Stuti Tyagi",
	}

	err := testStore.Create(context.TODO(), tempEmp)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	retrieved, err := testStore.GetAllEmployees(context.TODO(), PaginatedQuery{Limit: 10})
	if err != nil {
		t.Fatalf("failed to fetch employees: %v", err)
	}
	if len(retrieved) == 0 {
		t.Errorf("expected non-empty result, got empty slice")
	}
}

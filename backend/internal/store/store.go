package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// EmployeeStore interface with UpdateEmployee expecting an Employee struct
type EmployeeStore interface {
	Create(context.Context, Employee) error
	GetAllEmployees(context.Context) ([]*Employee, error)
	UpdateEmployee(context.Context, primitive.ObjectID, Employee) error // Changed to Employee
	DeleteEmployee(context.Context, primitive.ObjectID) error
}

// Storage store
type Storage struct {
	Employee EmployeeStore
}

// Mongo Implementation for EmployeeStore
type MongoEmployeeStore struct {
	collection *mongo.Collection
}

// Initializer function to initialize the storage with MongoDB client
func NewStorage(client *mongo.Client, dbName string) *Storage {
	db := client.Database(dbName)

	return &Storage{
		Employee: &MongoEmployeeStore{
			collection: db.Collection("stud"),
		},
	}
}

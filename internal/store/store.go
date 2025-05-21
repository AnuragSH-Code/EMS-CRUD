package store

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EmployeeStore interface {
	Create(context.Context, Employee) error
	GetAllEmployees(context.Context, PaginatedQuery) ([]*Employee, error)
	UpdateEmployee(context.Context, bson.ObjectID, bson.M) error
	DeleteEmployee(context.Context, bson.ObjectID) error
	IsEmailExists(context.Context, string) (bool, error)
}

type Storage struct {
	Employee EmployeeStore
}

type MongoEmployeeStore struct {
	collection *mongo.Collection
}

func NewStorage(client *mongo.Client, dbName string) *Storage {
	db := client.Database(dbName)

	return &Storage{
		Employee: &MongoEmployeeStore{
			collection: db.Collection("stud"),
		},
	}
}

package store

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockEmployeeStore struct {
	Employees   []*Employee
	ShouldError bool
	EmailExists bool
	UpdateError bool
	DeleteError bool
	LastUpdate  bson.M
}

func (m *MockEmployeeStore) GetAllEmployees(ctx context.Context, pq PaginatedQuery) ([]*Employee, error) {
	if m.ShouldError {
		return nil, errors.New("database error")
	}
	return m.Employees, nil
}

func (m *MockEmployeeStore) IsEmailExists(ctx context.Context, email string) (bool, error) {
	if m.ShouldError {
		return false, errors.New("database error")
	}
	return m.EmailExists, nil
}

func (m *MockEmployeeStore) Create(ctx context.Context, emp Employee) error {
	if m.ShouldError {
		return errors.New("database error")
	}
	m.Employees = append(m.Employees, &emp)
	return nil
}

func (m *MockEmployeeStore) UpdateEmployee(ctx context.Context, id bson.ObjectID, update bson.M) error {
	if m.UpdateError {
		return errors.New("update error")
	}
	m.LastUpdate = update
	return nil
}

func (m *MockEmployeeStore) DeleteEmployee(ctx context.Context, id bson.ObjectID) error {
	if m.DeleteError {
		return errors.New("delete error")
	}
	return nil
}

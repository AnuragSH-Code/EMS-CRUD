package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee struct with all the fields we need for an employee
type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string             `bson:"firstname" json:"firstname"`
	Lastname   string             `bson:"lastname" json:"lastname"`
	Role       string             `bson:"role" json:"role"`
	Department string             `bson:"department" json:"department"`
	Email      string             `bson:"email" json:"email"`
	ContactNo  string             `bson:"contact_no" json:"contact_no"`
	Manager    string             `bson:"manager" json:"manager"`
}

func (m *MongoEmployeeStore) Create(ctx context.Context, employee Employee) error {
	employee.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, employee)
	return err
}

// GetAllEmployees retrieves all employees
func (m *MongoEmployeeStore) GetAllEmployees(ctx context.Context) ([]*Employee, error) {
	var employees []*Employee

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var emp Employee
		if err := cursor.Decode(&emp); err != nil {
			return nil, err
		}
		employees = append(employees, &emp)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// GetEmployeeByID fetches an employee by their ID
func (m *MongoEmployeeStore) GetEmployeeByID(ctx context.Context, id primitive.ObjectID) (*Employee, error) {
	var employee Employee
	filter := bson.M{"_id": id}

	err := m.collection.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// UpdateEmployee updates employee information
func (m *MongoEmployeeStore) UpdateEmployee(ctx context.Context, id primitive.ObjectID, updatedEmployee Employee) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"firstname":  updatedEmployee.Firstname,
			"lastname":   updatedEmployee.Lastname,
			"role":       updatedEmployee.Role,
			"department": updatedEmployee.Department,
			"email":      updatedEmployee.Email,
			"contact_no": updatedEmployee.ContactNo,
			"manager":    updatedEmployee.Manager,
		},
	}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteEmployee deletes an employee by their ID
func (m *MongoEmployeeStore) DeleteEmployee(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := m.collection.DeleteOne(ctx, filter)
	return err
}

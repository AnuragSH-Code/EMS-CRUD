package store

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Employee struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string        `bson:"firstname" json:"firstname"`
	Lastname   string        `bson:"lastname" json:"lastname"`
	Role       string        `bson:"role" json:"role"`
	Department string        `bson:"department" json:"department"`
	Email      string        `bson:"email" json:"email"`
	ContactNo  string        `bson:"contact_no" json:"contact_no"`
	Manager    string        `bson:"manager" json:"manager"`
}

func (s *MongoEmployeeStore) IsEmailExists(ctx context.Context, email string) (bool, error) {
	filter := bson.M{"email": email}
	count, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *MongoEmployeeStore) Create(ctx context.Context, employee Employee) error {
	employee.ID = bson.NewObjectID()
	_, err := m.collection.InsertOne(ctx, employee)
	return err
}

func (m *MongoEmployeeStore) GetAllEmployees(ctx context.Context, pq PaginatedQuery) ([]*Employee, error) {
	var employees []*Employee

	opts := options.Find().
		SetLimit(int64(pq.Limit)).
		SetSkip(int64(pq.Offset))

	cursor, err := m.collection.Find(ctx, bson.M{}, opts)
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

func (m *MongoEmployeeStore) GetEmployeeByID(ctx context.Context, id bson.ObjectID) (*Employee, error) {
	var employee Employee
	filter := bson.M{"_id": id}

	err := m.collection.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (s *MongoEmployeeStore) UpdateEmployee(ctx context.Context, id bson.ObjectID, updates bson.M) error {
	filter := bson.M{"_id": id}
	updateQuery := bson.M{"$set": updates}
	_, err := s.collection.UpdateOne(ctx, filter, updateQuery)
	return err
}

func (m *MongoEmployeeStore) DeleteEmployee(ctx context.Context, id bson.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := m.collection.DeleteOne(ctx, filter)
	return err
}

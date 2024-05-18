package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const studentColl = "students"

type StudentStore interface {
	PostStudent(context.Context, *types.Student) (*types.Student, error)
}

type MongoStudentStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStudentStore(client *mongo.Client) *MongoStudentStore {
	return &MongoStudentStore{
		client: client,
		coll:   client.Database(DB_NAME).Collection(studentColl),
	}
}

func (s *MongoStudentStore) PostStudent(ctx context.Context, student *types.Student) (*types.Student, error) {
	cursor, err := s.coll.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}
	student.ID = cursor.InsertedID.(primitive.ObjectID)
	return student, nil
}

package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const studentColl = "students"

type StudentStore interface {
	GetStudents(ctx context.Context, filter bson.M) ([]*types.Student, error)
	PostStudent(ctx context.Context, student *types.Student) (*types.Student, error)
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

func (s *MongoStudentStore) GetStudents(ctx context.Context, filter bson.M) ([]*types.Student, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var students []*types.Student
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}
	return students, nil
}

func (s *MongoStudentStore) PostStudent(ctx context.Context, student *types.Student) (*types.Student, error) {
	cursor, err := s.coll.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}
	student.ID = cursor.InsertedID.(primitive.ObjectID)
	return student, nil
}

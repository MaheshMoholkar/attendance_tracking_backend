package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const attendanceColl = "attendance"

type AttendanceStore interface {
	GetAttendance(ctx context.Context, filter bson.M) ([]*types.Attendance, error)
	PostAttendance(ctx context.Context, attendance *types.Attendance) (*types.Attendance, error)
}

type MongoAttendanceStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoAttendanceStore(client *mongo.Client) *MongoAttendanceStore {
	return &MongoAttendanceStore{
		client: client,
		coll:   client.Database(DB_NAME).Collection(attendanceColl),
	}
}

func (s *MongoAttendanceStore) GetAttendance(ctx context.Context, filter bson.M) ([]*types.Attendance, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var attendance []*types.Attendance
	if err := cursor.All(ctx, &attendance); err != nil {
		return nil, err
	}
	return attendance, nil
}

func (s *MongoAttendanceStore) PostAttendance(ctx context.Context, attendance *types.Attendance) (*types.Attendance, error) {
	cursor, err := s.coll.InsertOne(ctx, attendance)
	if err != nil {
		return nil, err
	}
	attendance.ID = cursor.InsertedID.(primitive.ObjectID)
	return attendance, nil
}

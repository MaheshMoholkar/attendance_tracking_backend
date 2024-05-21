package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collegeColl = "college"

type CollegeStore interface {
	GetClasses(ctx context.Context, filter bson.M) ([]*types.Class, error)
}

type MongoCollegeStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCollegeStore(client *mongo.Client) *MongoCollegeStore {
	return &MongoCollegeStore{
		client: client,
		coll:   client.Database(DB_NAME).Collection(collegeColl),
	}
}

func (s *MongoCollegeStore) GetClasses(ctx context.Context, filter bson.M) ([]*types.Class, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var classes []*types.Class
	if err := cursor.All(ctx, &classes); err != nil {
		return nil, err
	}
	return classes, nil

}

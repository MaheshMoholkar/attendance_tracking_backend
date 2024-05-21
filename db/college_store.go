package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collegeColl = "college"

type CollegeStore interface {
	GetClassByName(ctx context.Context, filter bson.M) (*types.Class, error)
	GetClasses(ctx context.Context, filter bson.M) ([]*types.Class, error)
	GetClassInfo(ctx context.Context, filter bson.M) ([]*types.ClassInfo, error)
	PostClass(ctx context.Context, class *types.Class) (*types.Class, error)
	UpdateClass(ctx context.Context, filter, update bson.M) (*mongo.UpdateResult, error)
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

func (s *MongoCollegeStore) GetClassByName(ctx context.Context, filter bson.M) (*types.Class, error) {
	var class *types.Class
	err := s.coll.FindOne(ctx, filter).Decode(&class)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (s *MongoCollegeStore) GetClassInfo(ctx context.Context, filter bson.M) ([]*types.ClassInfo, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var classInfos []*types.ClassInfo
	for cursor.Next(ctx) {
		var class *types.Class
		if err := cursor.Decode(&class); err != nil {
			return nil, err
		}

		classInfo := &types.ClassInfo{
			ClassName: class.ClassName,
			Divisions: make(map[string]bool),
		}

		// Extract division names without student IDs
		for division := range class.Divisions {
			classInfo.Divisions[division] = true
		}

		classInfos = append(classInfos, classInfo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return classInfos, nil
}

func (s *MongoCollegeStore) GetClasses(ctx context.Context, filter bson.M) ([]*types.Class, error) {
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var classes []*types.Class
	if err := cursor.All(ctx, &classes); err != nil {
		return nil, err
	}
	return classes, nil

}

func (s *MongoCollegeStore) PostClass(ctx context.Context, class *types.Class) (*types.Class, error) {
	cursor, err := s.coll.InsertOne(ctx, class)
	if err != nil {
		return nil, err
	}
	class.ID = cursor.InsertedID.(primitive.ObjectID)
	return class, nil
}

func (s *MongoCollegeStore) UpdateClass(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return s.coll.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
}

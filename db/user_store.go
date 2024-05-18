package db

import (
	"context"

	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUserByEmail(context.Context, string) (string, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DB_NAME).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (string, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	cursor, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = cursor.InsertedID.(primitive.ObjectID)
	return user, nil
}

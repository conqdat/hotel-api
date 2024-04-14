package db

import (
	"context"
	"github.com/conqdat/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME = "hotel-reservation"
const userColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoDBStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoDBStore {
	return &MongoDBStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoDBStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": ToObjectID(id)}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

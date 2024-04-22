package db

import (
	"context"
	"fmt"
	"github.com/conqdat/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection("hotels"),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("error: inserted ID is not a valid ObjectID")
	}
	hotel.ID = insertedID
	return hotel, nil
}

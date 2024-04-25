package main

import (
	"context"
	"fmt"
	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Size:  "small",
			Price: 1999.9,
		},
		{
			Size:  "normal",
			Price: 1999.9,
		},
		{
			Size:  "small",
			Price: 1999.9,
		},
		{
			Size:  "medium",
			Price: 1999.9,
		},
	}
	_, err := hotelStore.InsertHotel(context.Background(), &hotel)
	if err != nil {
		log.Fatalln(err)
	}

	for _, room := range rooms {
		room.HotelID = hotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Printf("seed %v hotel successfully \n", name)
}

func main() {
	seedHotel("Bellucia", "US", 3)
	seedHotel("VinFast", "UK", 4)
	seedHotel("Something", "JP", 5)
}

func init() {
	var err error
	ctx := context.Background()
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatalln(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatalln(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}

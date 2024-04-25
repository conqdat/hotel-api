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

func seedHotel(name, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 1999.9,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      types.SeaSideRoomType,
			BasePrice: 2399.9,
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
	seedHotel("Bellucia", "US")
	seedHotel("VinFast", "UK")
	seedHotel("Something", "JP")
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

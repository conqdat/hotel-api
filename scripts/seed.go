package main

import (
	"context"
	"fmt"
	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatalln(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

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

	hotel := types.Hotel{
		Name:     "Hotel one",
		Location: "US",
		Rooms:    nil,
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatalln(err)
	}

	for _, room := range rooms {
		room.HotelID = hotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(insertedRoom)
	}
	fmt.Println(insertedHotel)
}

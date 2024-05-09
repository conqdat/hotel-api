package main

import (
	"context"
	"fmt"
	"github.com/conqdat/hotel-api/api"
	"log"
	"strconv"

	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	userStore  db.UserStore
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

func seedUser(isAdmin bool, firstName, lastName, email string) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  "12345",
	})
	if err != nil {
		log.Fatalln(err)
	}
	user.IsAdmin = isAdmin
	user, err = userStore.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("seed %v %v user successfully \n", user.FirstName, user.LastName)
	fmt.Println("token: ", api.CreateTokenFromUser(user))
}

func main() {

	for i := 0; i < 100; i++ {
		seedHotel("Bellucia "+strconv.Itoa(i), "US", 3)
		seedHotel("VinFast "+strconv.Itoa(i), "UK", 4)
		seedHotel("Something "+strconv.Itoa(i), "JP", 5)
	}

	seedUser(true, "Dat 1", "Tran", "trandat1@gmail.com")
	seedUser(false, "Dat 2", "Tran", "trandat2@gmail.com")
	seedUser(false, "Dat 3", "Tran", "trandat3@gmail.com")
	seedUser(false, "Dat 4", "Tran", "trandat4@gmail.com")
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
	userStore = db.NewMongoUserStore(client)
}

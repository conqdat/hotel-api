package main

import (
	"context"
	"log"

	"github.com/conqdat/hotel-api/api"
	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatalln(err)
	}

	var (

		// Init Handler
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userStore  	 = db.NewMongoUserStore(client)
		bookingStore = db.NewBookingStore(client)
		
		store      = &db.Store{
			Hotel: hotelStore,
			Room:  roomStore,
			User:  userStore,
			Booking: bookingStore,
		}

		hotelHandler = api.NewHotelHandler(store)
		userHandler  = api.NewUserHandler(userStore)
		authHandler  = api.NewAuthHandle(userStore)
		roomHandler  = api.NewRoomHandler(store)

		app     = fiber.New(config)
		apiV1   = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
		apiAuth = app.Group("/api")
	)

	// User Handlers
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandleCreateUser)

	// Authentication Handlers
	apiAuth.Post("/login", authHandler.HandleAuthenticate)

	// Hotel Handlers
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotels/:id/rooms", hotelHandler.HandleGetRoom)
	apiV1.Get("/hotels/:id", hotelHandler.HandleGetHotelByID)

	// Room handlers
	apiV1.Post("/room/:id/book", roomHandler.HandleBookRoom)
	apiV1.Get("/bookings", roomHandler.HandleGetBookings)
	apiV1.Get("/room", roomHandler.HandleGetRooms)

	app.Get("/", handleHelloWord)
	app.Listen(":3000")
}

func handleHelloWord(c *fiber.Ctx) error {
	return c.JSON("Hello guys !!!")
}

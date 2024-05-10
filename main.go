package main

import (
	"context"
	"github.com/conqdat/hotel-api/api"
	"github.com/conqdat/hotel-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(api.Error); ok {
			return c.Status(apiError.Code).JSON(apiError)
		}
		return api.NewError(http.StatusInternalServerError, err.Error())
	},
}

func init() {

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
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewBookingStore(client)

		store = &db.Store{
			Hotel:   hotelStore,
			Room:    roomStore,
			User:    userStore,
			Booking: bookingStore,
		}

		hotelHandler   = api.NewHotelHandler(store)
		userHandler    = api.NewUserHandler(userStore)
		authHandler    = api.NewAuthHandle(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)

		app     = fiber.New(config)
		apiV1   = app.Group("/api/v1", api.JWTAuthentication(userStore))
		apiAuth = app.Group("/api")
		admin   = apiV1.Group("/admin", api.AdminAuth)
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

	// Booking Handlers
	// ADMIN
	admin.Get("/bookings", bookingHandler.HandleGetBookings)

	apiV1.Get("/bookings/:id", bookingHandler.HandleGetBooking)
	apiV1.Get("/bookings/:id/cancel", bookingHandler.HandleCancelBooking)

	app.Get("/", handleHelloWord)
	app.Listen(":3000")
}

func handleHelloWord(c *fiber.Ctx) error {
	return c.JSON("Hello guys !!!")
}

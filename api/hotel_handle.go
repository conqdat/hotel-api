package api

import (
	"github.com/conqdat/hotel-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type RecourseResponse struct {
	Total int `json:"total"`
	Data  any `json:"data"`
	Page  int `json:"page"`
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1"))
	rating, _ := strconv.Atoi(c.Query("rating", "1"))

	filter := bson.M{
		"rating": rating,
	}

	opts := options.FindOptions{}
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetLimit(int64(limit))
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &opts)
	if err != nil {
		return err
	}
	res := RecourseResponse{
		Total: len(hotels),
		Data:  hotels,
		Page:  page,
	}
	return c.JSON(res)
}

func (h *HotelHandler) HandleGetRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return InvalidID()
	}
	filter := bson.M{"hotelid": oid} // I dont know why hotelID not work !!!
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return InvalidID()
	}
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}

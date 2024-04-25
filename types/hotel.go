package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // mongo DB => bson
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // mongo DB => bson
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"Price" json:"Price"`
	HotelID   primitive.ObjectID `json:"hotelID" json:"hotelID"`
}

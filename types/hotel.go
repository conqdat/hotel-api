package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // mongo DB => bson
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}


type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // mongo DB => bson
	Size    string             `bson:"size" json:"size"`                  // small / medium / large
	Seaside bool               `bson:"seaside" json:"seaside"`
	Price   float64            `bson:"Price" json:"Price"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

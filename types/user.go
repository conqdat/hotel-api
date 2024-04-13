package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"` // mongo DB => bson
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}

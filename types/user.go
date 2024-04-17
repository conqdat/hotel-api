package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost         = 12
	minLengthFirstName = 3
	minLengthLastName  = 3
	minPasswordLength  = 3
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	var errors []string
	if len(params.Password) < minPasswordLength {
		errors = append(errors, fmt.Sprintf("password length should be at least %d characters", minPasswordLength))
	}
	if len(params.FirstName) < minLengthFirstName {
		errors = append(errors, fmt.Sprintf("firstName length should be at least %d characters", minLengthFirstName))
	}
	if len(params.LastName) < minLengthLastName {
		errors = append(errors, fmt.Sprintf("lastName length should be at least %d characters", minLengthLastName))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Sprintf("email is valid"))
	}
	return errors
}

func isEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // mongo DB => bson
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"-" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	enpwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(enpwd),
	}, nil
}

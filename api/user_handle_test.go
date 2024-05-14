package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http/httptest"
	"testing"
)

const testdburi = "mongodb://localhost:27017"
const testdbname = "hotel-reservation-test"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatalln(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		FirstName: "test first name",
		LastName:  "test last name",
		Email:     "test email",
		Password:  "12345",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	var user types.User
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return
	}
	if len(user.ID) == 0 {
		t.Errorf("expected a user id to be set")
	}
	if len(user.EncryptedPassword) == 0 {
		t.Errorf("expected the EncryptedPass not to be included in the json response")
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
}

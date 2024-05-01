package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"github.com/gofiber/fiber/v2"
)


func insertTestUser(_ *testing.T, userStore db.UserStore) (*types.User, error) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: "firstName",
		LastName:  "lastName",
		Email:     "trandat1@gmail.com",
		Password:  "12345",
	})
	if err != nil {
		log.Fatalln(err)
	}
	insertedUser, err := userStore.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser, nil
}

func TestAuthenticateWithWrongPassWordFailure(t *testing.T) {
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertedUser, err := insertTestUser(t, tdb.UserStore)
	if err != nil {
		t.Fatal(err)
	}

	app := fiber.New()
	authHandler := NewAuthHandle(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email: "trandat1@gmail.com",
		Password: "12345",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected http status code 200 but got %d", res.StatusCode)
	}
	var authRes AuthResponse 
	if err := json.NewDecoder(res.Body).Decode(&authRes); err != nil {
		t.Fatal(err)
	}
	if len(authRes.Token) == 0 {
		t.Fatal("expected JWT token to be present in auth response")
	}
	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insertedUser, authRes.User) {
		t.Fatal("expected the same user")
	}
}
package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters.")
	}

	// Does user with this username already exist?
	exists, err := api.dbStore.UserExists(event.Username)
	if err != nil {
		return fmt.Errorf("There was an error checking if the user exists. %w", err)
	}

	if exists {
		return fmt.Errorf("A user with that username already exists.")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("There was an error registering the user. %w", err)
	}
	return nil
}

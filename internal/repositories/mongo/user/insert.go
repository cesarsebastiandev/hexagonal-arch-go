package user

import (
	"context"
	"fmt"
	"hexagonal-architecture/internal/domain"
	"log"
)

func (r Repository) Insert(user domain.User) (id interface{}, err error) {

	collection := r.Client.Database("dbusers").Collection("users")
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return insertResult.InsertedID, nil
}

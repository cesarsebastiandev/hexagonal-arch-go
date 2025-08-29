package user

import (
	"fmt"
	"hexagonal-architecture/internal/domain"
	"log"
	"time"
)

func (s Service) Create(user domain.User) (id interface{}, err error) {

	//Set creation time
	//Save repo
	//Respond with the id and resource created
	user.CreationTime = time.Now().UTC()
	//REPO
	insertedId, err := s.Repo.Insert(user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	//REPO

	return insertedId, nil

}

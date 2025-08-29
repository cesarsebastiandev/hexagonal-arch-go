package ports

import "hexagonal-architecture/internal/domain"

type UserService interface {
	Create(user domain.User) (id any, err error)
}

type UserRepository interface {
	Insert(user domain.User) (id any, err error)

}


package user

import "hexagonal-architecture/internal/ports"

type Handler struct {
	UserService ports.UserService
}


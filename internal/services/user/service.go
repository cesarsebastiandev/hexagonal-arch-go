package user

import "hexagonal-architecture/internal/ports"

type Service struct {
	Repo ports.UserRepository
}

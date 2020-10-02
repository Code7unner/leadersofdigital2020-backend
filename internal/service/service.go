package service

import "github.com/code7unner/leadersofdigital2020-backend/internal/repository"

type Servicer interface {
	User() UserServicer
}

type Service struct {
	/* Репозиторий для работы с базой */
	repository repository.Repository

	user *UserService
}

type UserServicer interface {
	GetUser(string) error
}

func New(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) User() UserServicer {
	if s.user != nil {
		return s.user
	}

	s.user = &UserService{
		service: s,
	}

	return s.user
}

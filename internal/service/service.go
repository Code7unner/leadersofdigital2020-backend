package service

import "github.com/code7unner/leadersofdigital2020-backend/internal/repository"

type Servicer interface {
	Test() Tester
}

type Service struct {
	/* Репозиторий для работы с базой */
	repository repository.Repository

	test       *TestService
}

type Tester interface {
	PrintTest(string) error
}

func New(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Test() Tester {
	if s.test != nil {
		return s.test
	}

	s.test = &TestService{
		service: s,
	}

	return s.test
}
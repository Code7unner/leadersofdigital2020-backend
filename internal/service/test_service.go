package service

type TestService struct {
	service *Service
}

func (s *TestService) PrintTest(string) error {
	panic("")
}
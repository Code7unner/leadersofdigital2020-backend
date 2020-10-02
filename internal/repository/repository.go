package repository

type Repository interface {
	Test() TestRepository
}

type TestRepository interface {
	Print(string) error
}
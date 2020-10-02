package repository

type Repository interface {
	User() UserRepository
	Products() ProductsRepository
}

type UserRepository interface {
	Print(string) error
}

type ProductsRepository interface {
	Get(string) error
}
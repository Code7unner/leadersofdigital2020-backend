package repository_implementation

type ProductsRepository struct {
	repository *Repository
}

func (t ProductsRepository) Get(s string) error {
	panic("implement me")
}

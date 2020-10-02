package repository_implementation

type TestRepository struct {
	repository *Repository
}

func (t TestRepository) Print(s string) error {
	panic("implement me")
}


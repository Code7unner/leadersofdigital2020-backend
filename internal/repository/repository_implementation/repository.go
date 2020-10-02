package repository_implementation

import (
	"database/sql"
	"github.com/code7unner/leadersofdigital2020-backend/internal/repository"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

var (
	dialect = postgresql.Dialect
)

type Repository struct {
	db       *reform.DB
	user     *UserRepository
	products *ProductsRepository
}

func (r *Repository) Products() repository.ProductsRepository {
	if r.products != nil {
		return r.products
	}

	r.user = &UserRepository{
		repository: r,
	}
	return r.products
}

func New(db *sql.DB, printf repository.Printf) repository.Repository {
	repo := new(Repository)

	logger := reform.NewPrintfLogger(reform.Printf(printf))
	repo.db = reform.NewDB(db, dialect, logger)

	return repo
}

func (r *Repository) User() repository.UserRepository {
	if r.user != nil {
		return r.user
	}

	r.user = &UserRepository{
		repository: r,
	}
	return r.user
}

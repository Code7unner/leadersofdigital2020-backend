package repository_implementation

import (
	"database/sql"
	"github.com/code7unner/leadersofdigital2020-backend/internal/repository"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	//"database/sql"
)

var (
	dialect = postgresql.Dialect
)

type Repository struct {
	db          *reform.DB
	test *TestRepository
}

func (r *Repository) Test() repository.TestRepository {
	if r.test != nil {
		return r.test
	}

	r.test = &TestRepository{
		repository: r,
	}
	return r.test
}

func New(db *sql.DB, printf repository.Printf) repository.Repository {
	repo := new(Repository)

	logger := reform.NewPrintfLogger(reform.Printf(printf))
	repo.db = reform.NewDB(db, dialect, logger)

	return repo
}


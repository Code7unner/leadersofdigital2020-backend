package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/interrupt"
	"github.com/code7unner/leadersofdigital2020-backend/internal/logging"
	"github.com/code7unner/leadersofdigital2020-backend/internal/repository/repository_implementation"
	"github.com/code7unner/leadersofdigital2020-backend/internal/routes"
	"github.com/code7unner/leadersofdigital2020-backend/internal/server"
	newService "github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"log"
)

var (
	envPathFlag string
)

func init() {
	flag.StringVar(&envPathFlag, "env", ".env", "env file path")
	flag.Parse()
}

func main() {
	ctx, done := interrupt.Context()
	defer done()

	// Initialize config
	conf := configs.GetCommonEnvConfigs()

	if err := runExternalServer(ctx, &conf); err != nil {
		logger := logging.FromContext(ctx)
		logger.Fatal(err)
	}
}

func runExternalServer(ctx context.Context, conf *configs.CommonEnvConfigs) error {
	logger := logging.FromContext(ctx)

	conn, err := newDB(conf.PostgresDBStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repository := repository_implementation.New(conn,
		logger.Infof,
	)

	service := newService.New(repository)

	externalRouter := chi.NewRouter()
	// Test requests
	externalRouter.Group(routes.TestRoutes(service))

	srv, err := server.New(conf.ServerExternalPort)
	if err != nil {
		return fmt.Errorf("server.New: %w", err)
	}
	logger.Infof("listening on :%s", conf.ServerExternalPort)
	return srv.ServeHTTPHandler(ctx, externalRouter)
}

func newDB(sqlConnString string) (*sql.DB, error) {
	pgxConf, err := pgx.ParseConfig(sqlConnString)
	if err != nil {
		return nil, err
	}

	conn := stdlib.OpenDB(*pgxConf)

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(3)
	return conn, nil
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/internal/interrupt"
	"github.com/code7unner/leadersofdigital2020-backend/internal/logging"
	"github.com/code7unner/leadersofdigital2020-backend/internal/routes"
	"github.com/code7unner/leadersofdigital2020-backend/internal/server"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
)

func main() {
	ctx, done := interrupt.Context()
	defer done()

	// Initialize logger
	logger := logging.FromContext(ctx)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Initialize config
	config := configs.NewConfig()

	if err := runExternalServer(ctx, config, logger); err != nil {
		logger := logging.FromContext(ctx)
		logger.Fatal(err)
	}
}

func runExternalServer(ctx context.Context, conf *configs.Config, logger *zap.SugaredLogger) error {
	conn, err := newDB(conf.PostgresDBStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Init db service
	storeStorage := db.NewStoreStorage(conn)

	externalRouter := chi.NewRouter()
	// Test requests
	externalRouter.Group(routes.InitRoutes(storeStorage))

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

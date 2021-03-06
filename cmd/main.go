package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/auth"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/internal/interrupt"
	"github.com/code7unner/leadersofdigital2020-backend/internal/logging"
	"github.com/code7unner/leadersofdigital2020-backend/internal/routes"
	"github.com/code7unner/leadersofdigital2020-backend/internal/server"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
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

func runExternalServer(ctx context.Context, config *configs.Config, logger *zap.SugaredLogger) error {
	var conn *sql.DB
	if config.AppEnv == "production" {
		conn, err := newDB(config.PostgresDBStr)
		if err != nil {
			return err
		}
		defer conn.Close()
	} else {
		conn, err := newDB(config.PostgresTestDBStr)
		if err != nil {
			return err
		}
		defer conn.Close()
	}

	// Init db storages
	var storages = []db.Storage{
		db.NewUserStorage(conn),
		db.NewProductStorage(conn),
		db.NewOrderStorage(conn),
		db.NewStoreStorage(conn),
	}

	r := chi.NewRouter()

	// Protected router
	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(auth.New("HS256", []byte(config.TokenSecret), nil)))
		r.Use(auth.Authenticator)

		r.Route("/api/v1", routes.InitRoutes(config, storages...))
	})

	root := "./dist"
	fs := http.FileServer(http.Dir(root))

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
				http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
			} else {
				fs.ServeHTTP(w, r)
			}
		})

		r.Post("/register", auth.Register(config.TokenSecret))
	})

	// Cors
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler(r)

	srv, err := server.New(config.ServerExternalPort)
	if err != nil {
		return fmt.Errorf("server.New: %w", err)
	}
	logger.Infof("listening on :%s", config.ServerExternalPort)
	return srv.ServeHTTPHandler(ctx, handler)
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

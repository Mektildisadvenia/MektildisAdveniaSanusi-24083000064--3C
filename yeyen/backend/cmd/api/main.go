package main

import (
	"backend/config"
	"backend/internal/handler/http"
	"backend/internal/repository/postgres"
	"backend/internal/usecase"
	"backend/pkg/database"
	"log"
	nethttp "net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg.DBURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Repositories
	userRepo := postgres.NewUserRepository(db)
	productRepo := postgres.NewProductRepository(db)

	// Usecases
	userUsecase := usecase.NewUserUsecase(userRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	// Handlers
	userHandler := http.NewUserHandler(userUsecase)
	productHandler := http.NewProductHandler(productUsecase)

	mux := nethttp.NewServeMux()

	// Routes
	userHandler.RegisterRoutes(mux)
	productHandler.RegisterRoutes(mux)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := nethttp.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

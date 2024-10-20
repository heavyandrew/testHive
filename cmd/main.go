// @title testHive API
// @version 1.0
// @description This is a testHive API realization.

// @host localhost:8080
// @BasePath /api
// @schemes http
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testHive/database"
	"testHive/internal/config"
	"testHive/internal/handlers"
	"testHive/internal/repository"
	"testHive/internal/services"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "testHive/docs"
)

func main() {
	cfg := config.GetConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	assetRepo := repository.NewAssetRepository(db)

	userService := services.NewUserService(userRepo)
	assetService := services.NewAssetService(assetRepo)

	userHandler := handlers.NewUserHandler(userService, cfg.JWTSecret)
	assetHandler := handlers.NewAssetHandler(assetService, userHandler)

	r := mux.NewRouter()

	r.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")

	r.HandleFunc("/api/assets", assetHandler.AddAsset).Methods("POST")
	r.HandleFunc("/api/assets", assetHandler.SearchAssets).Methods("GET")
	r.HandleFunc("/api/assets/{id}", assetHandler.DeleteAsset).Methods("DELETE")

	r.HandleFunc("/api/assets/buy", assetHandler.BuyAsset).Methods("POST")
	r.HandleFunc("/api/assets/my", assetHandler.GetUserAssets).Methods("GET")

	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	log.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", r)
}

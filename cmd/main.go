package main

import (
	"log"
	"net/http"
	"testHive/database"
	"testHive/internal/config"
	"testHive/internal/handlers"
	"testHive/internal/repository"
	"testHive/internal/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	// _ "testHive/docs" // подключение автоматически сгенерированной документации Swagger
)

// @title MyApp API
// @version 1.0
// @description Документация для MyApp API
// @host localhost:8080
// @BasePath /api

func main() {
	// Загружаем конфигурацию
	cfg := config.GetConfig()

	// Подключение к базе данных
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Создаем репозитории
	userRepo := repository.NewUserRepository(db)
	assetRepo := repository.NewAssetRepository(db)

	// Создаем сервисы
	userService := services.NewUserService(userRepo)
	assetService := services.NewAssetService(assetRepo)

	// Создаем обработчики
	userHandler := handlers.NewUserHandler(userService, cfg.JWTSecret)
	assetHandler := handlers.NewAssetHandler(assetService, userHandler)

	// Настройка маршрутов
	r := mux.NewRouter()

	// Маршруты для пользователей
	r.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")

	// Маршруты для ассетов
	r.HandleFunc("/api/assets", assetHandler.AddAsset).Methods("POST")
	r.HandleFunc("/api/assets", assetHandler.GetUserAssets).Methods("GET")
	r.HandleFunc("/api/assets/{id}", assetHandler.DeleteAsset).Methods("DELETE")

	// Swagger-документация
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Запуск сервера
	log.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", r)
}

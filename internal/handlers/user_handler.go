package handlers

import (
	"encoding/json"
	"net/http"
	"testHive/internal/auth"
	"testHive/internal/models"
	"testHive/internal/services"
	"time"
)

type UserHandler struct {
	UserService *services.UserService
	JWTSecret   string
}

func NewUserHandler(userService *services.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{UserService: userService, JWTSecret: jwtSecret}
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Создание нового пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "User"
// @Success 201
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Could not create user"
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	if err := h.UserService.RegisterUser(&user); err != nil {
		http.Error(w, "Ошибка при регистрации", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Summary Вход пользователя
// @Description Аутентификация пользователя и генерация JWT токена
// @Tags users
// @Accept  json
// @Produce  json
// @Param   credentials  body  models.User  true  "User Credentials"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Invalid username or password"
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	user, err := h.UserService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}
	// Генерация JWT токена с использованием секретного ключа
	token, err := auth.GenerateJWT(user.ID, h.JWTSecret, time.Hour*24)
	if err != nil {
		http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

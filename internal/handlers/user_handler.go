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
	token, err := auth.GenerateJWT(user.ID, h.JWTSecret, time.Hour*24)
	if err != nil {
		http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (h *UserHandler) Authorize(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, http.ErrNoCookie
	}

	tokenStr := authHeader[len("Bearer "):]

	claims, err := auth.ValidateJWT(tokenStr, h.JWTSecret)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}

// Register - godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	username, err := h.UserService.UserAlreadyExists(&user)
	if err != nil {
		http.Error(w, "Не смогли проверить, что пользователь уже существует", http.StatusInternalServerError)
		return
	}
	if username != nil {
		http.Error(w, "Пользователь с таким ником уже существует", http.StatusBadRequest)
		return
	}
	if err := h.UserService.RegisterUser(&user); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при регистрации: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Login - godoc
// @Summary Login a user
// @Description Login a user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/login [post]
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
	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func (h *UserHandler) getUserIDFromToken(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("заголовок Authorization отсутствует")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, fmt.Errorf("неправильный формат заголовка Authorization")
	}

	tokenStr := parts[1]

	claims, err := auth.ValidateJWT(tokenStr, h.JWTSecret)
	if err != nil {
		return 0, fmt.Errorf("ошибка валидации токена: %v", err)
	}

	return claims.UserID, nil
}

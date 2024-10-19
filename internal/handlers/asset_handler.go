package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testHive/internal/models"
	"testHive/internal/services"
)

type AssetHandler struct {
	AssetService *services.AssetService
	UserHandler  *UserHandler
}

func NewAssetHandler(assetService *services.AssetService, userHandler *UserHandler) *AssetHandler {
	return &AssetHandler{AssetService: assetService, UserHandler: userHandler}
}

func (h *AssetHandler) AddAsset(w http.ResponseWriter, r *http.Request) {
	userID, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	if err := h.AssetService.AddAsset(userID, &asset); err != nil {
		http.Error(w, "Ошибка при добавлении актива", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AssetHandler) GetUserAssets(w http.ResponseWriter, r *http.Request) {
	userID, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}
	assets, err := h.AssetService.GetUserAssets(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении активов", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(assets)
}

func (h *AssetHandler) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	assetID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Неверный идентификатор актива", http.StatusBadRequest)
		return
	}
	if err := h.AssetService.DeleteAsset(assetID); err != nil {
		http.Error(w, "Ошибка при удалении актива", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

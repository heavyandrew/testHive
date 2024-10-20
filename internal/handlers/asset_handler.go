package handlers

import (
	"encoding/json"
	"fmt"
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

type BuyAssetRequest struct {
	AssetID int     `json:"asset_id"`
	Price   float64 `json:"price"`
}

func NewAssetHandler(assetService *services.AssetService, userHandler *UserHandler) *AssetHandler {
	return &AssetHandler{AssetService: assetService, UserHandler: userHandler}
}

// AddAsset - godoc
// @Summary Add a new asset
// @Description Add a new asset with details
// @Tags assets
// @Accept json
// @Produce json
// @Param asset body models.Asset true "Asset data"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/assets [post]
func (h *AssetHandler) AddAsset(w http.ResponseWriter, r *http.Request) {
	_, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	assetId, err := h.AssetService.AddAsset(&asset)
	if err != nil {
		http.Error(w, "Ошибка при добавлении актива", http.StatusInternalServerError)
		return
	}

	response := struct {
		AssetId int `json:"id"`
	}{AssetId: assetId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SearchAssets - godoc
// @Summary Search for assets
// @Description Search assets by query string
// @Tags assets
// @Accept json
// @Produce json
// @Param search query string false "Search query"
// @Success 200 {array} models.Asset
// @Failure 500 {object} map[string]interface{}
// @Router /api/assets [get]
func (h *AssetHandler) SearchAssets(w http.ResponseWriter, r *http.Request) {
	_, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}
	search := r.URL.Query().Get("search")

	assets, err := h.AssetService.FindAssets(search)
	if err != nil {
		http.Error(w, "Ошибка при поиске активов", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assets)
}

// DeleteAsset - godoc
// @Summary Delete an asset
// @Description Delete an asset by ID
// @Tags assets
// @Accept json
// @Produce json
// @Param id path int true "Asset ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/assets/{id} [delete]
func (h *AssetHandler) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	_, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

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

// BuyAsset - godoc
// @Summary Buy an asset
// @Description Buy an asset by ID with price
// @Tags assets
// @Accept json
// @Produce json
// @Param asset body BuyAssetRequest true "Asset purchase data"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/assets/buy [post]
func (h *AssetHandler) BuyAsset(w http.ResponseWriter, r *http.Request) {
	_, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

	var req struct {
		AssetID int     `json:"asset_id"`
		Price   float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	userID, err := h.UserHandler.getUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Не удалось определить Id пользователя", http.StatusInternalServerError)
		return
	}

	if err := h.AssetService.BuyAsset(userID, req.AssetID, req.Price); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при покупке актива: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUserAssets - godoc
// @Summary Get user's assets
// @Description Get all assets belonging to a user
// @Tags assets
// @Accept json
// @Produce json
// @Success 200 {array} models.Asset
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/assets/my [get]
func (h *AssetHandler) GetUserAssets(w http.ResponseWriter, r *http.Request) {
	_, err := h.UserHandler.Authorize(r)
	if err != nil {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

	userID, err := h.UserHandler.getUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Не удалось определить Id пользователя", http.StatusInternalServerError)
		return
	}

	assets, err := h.AssetService.GetUserAssets(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении активов пользователя", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(assets)
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/usecase"
)

// ============================================================================
// HANDLER LAYER - HTTP handlers (presentación)
// ============================================================================
// Los handlers son responsables de:
// - Parsear requests HTTP
// - Validar entrada
// - Llamar a los casos de uso
// - Formatear respuestas
// ============================================================================

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.CreateUser(r.Context(), req.Email, req.Name, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	})
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	})
}


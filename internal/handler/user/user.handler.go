package user_handler

import (
	"encoding/json"
	"net/http"
	"time"

	user_service "github.com/098765432m/monthly_planner_backend/internal/service/user"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type UserHandler struct {
	service *user_service.UserService
}

func NewUserHandler(service *user_service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateUser)
	r.Get("/{id}", h.GetUserById)
	r.Put("/{id}", h.UpdateUserById)
	r.Delete("/{id}", h.DeleteUserById)

	return r
}

type UserResponse struct {
	ID          string    `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	IsActive    bool      `json:"is_active,omitempty"`
	Role        string    `json:"role,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// TODO Add Login, Register

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	user, err := h.service.GetUserById(r.Context(), idParam)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get user")
	}

	userResp := UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
	}

	utils.WriteJSON(w, http.StatusOK, userResp, "")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req user_service.CreateUserServiceParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Bad Request Error")
		return
	}

	user, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		zap.S().Error(err)
		utils.WriteError(w, http.StatusBadRequest, "Failed to get User")
		return
	}

	userResp := UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	utils.WriteJSON(w, http.StatusCreated, userResp, "Tao tai khoan thanh cong!")
}

func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var req user_service.UpdateUserByIdServiceParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid Request params")
		return
	}

	err := h.service.UpdateUserById(r.Context(), idParam, &user_service.UpdateUserByIdServiceParams{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
		IsActive:    req.IsActive,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to Update user")
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil, "Cap nhat nguoi dung thanh cong!")
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	if err := h.service.DeleteUserById(r.Context(), idParam); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to Delete User")
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil, "Xoa tai khoan thanh cong!")
}

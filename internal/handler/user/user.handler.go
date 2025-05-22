package user_handler

import (
	"net/http"
	"time"

	user_service "github.com/098765432m/monthly_planner_backend/internal/service/user"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user_service.UserService
}

func NewUserHandler(service *user_service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")

	user.POST("/", h.CreateUser)
	user.GET("/:id", h.GetUserById)
	user.PUT("/:id", h.UpdateUserById)
	user.DELETE("/:id", h.DeleteUserById)

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

func (h *UserHandler) GetUserById(c *gin.Context) {
	idParam := c.Param("id")

	user, err := h.service.GetUserById(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Get User!"))
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

	c.JSON(http.StatusCreated, utils.SuccessResponse(userResp, ""))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req user_service.CreateUserServiceParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Bad Request Error"))
	}

	user, err := h.service.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Create User!"))
	}

	userResp := UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(userResp, "Tao tai khoan thanh cong!"))
}

func (h *UserHandler) UpdateUserById(c *gin.Context) {
	idParam := c.Param("id")

	var req user_service.UpdateUserByIdServiceParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Bad Request Error"))
		return
	}

	err := h.service.UpdateUserById(c.Request.Context(), idParam, &user_service.UpdateUserByIdServiceParams{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
		IsActive:    req.IsActive,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Update User!"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Cap nhat tai khoan thanh cong!"))
}

func (h *UserHandler) DeleteUserById(c *gin.Context) {
	idParam := c.Param("id")

	if err := h.service.DeleteUserById(c.Request.Context(), idParam); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Bad Request Error"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Xoa tai khoan thanh cong!"))
}

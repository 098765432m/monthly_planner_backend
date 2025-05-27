package task_category_handler

import (
	"net/http"

	task_category_service "github.com/098765432m/monthly_planner_backend/internal/service/task_category"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskCategoryHandler struct {
	service *task_category_service.TaskCategoryService
}

func NewTaskCategoryHandler(
	service *task_category_service.TaskCategoryService) *TaskCategoryHandler {
	return &TaskCategoryHandler{
		service: service,
	}
}

func (h *TaskCategoryHandler) RegisterRoutes(rg *gin.RouterGroup) {
	task_category := rg.Group("/taskCategories")

	task_category.POST("/", h.CreateTaskCategory)
	task_category.DELETE("/:id", h.DeleteTaskCategory)
}

func (h *TaskCategoryHandler) CreateTaskCategory(c *gin.Context) {
	var req task_category_service.CreateTaskCategoryServiceParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to parse body!"))
		return
	}

	resultId, err := h.service.CreateTaskCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Create Task Category!"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(resultId, ""))
}

func (h *TaskCategoryHandler) DeleteTaskCategory(c *gin.Context) {
	idParam := c.Param("id")
	var id pgtype.UUID
	if err := id.Scan(idParam); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid UUID!"))
		return
	}

	if err := h.service.DeleteTaskCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Detele Task Category!"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Xoa thanh cong!"))
}

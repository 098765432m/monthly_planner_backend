package month_handler

import (
	"net/http"
	"time"

	month_service "github.com/098765432m/monthly_planner_backend/internal/service/month"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type MonthHandler struct {
	service *month_service.MonthService
}

func NewMonthHandler(service *month_service.MonthService) *MonthHandler {
	return &MonthHandler{
		service: service,
	}
}

func (h *MonthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	month := rg.Group("/months")

	month.POST("/", h.CreateMonth)
	month.DELETE("/:id", h.DeleteMonth)
}

func (h *MonthHandler) Test(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	zap.S().Infoln(idStr)

	sizeStr := r.URL.Query().Get("size")
	zap.S().Infoln("size: ", sizeStr)

	// err := json.NewDecoder().Decode()
}

type CreateMonthRequest struct {
	Month string `json:"month"`
}

func (h *MonthHandler) CreateMonth(c *gin.Context) {
	var req CreateMonthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Month Date"))
		return
	}

	parsedTime, err := time.Parse("01/2006", req.Month)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invadlid Month Date")
		return
	}

	err = h.service.CreateMonth(c.Copy().Request.Context(), int8(parsedTime.Month()), int16(parsedTime.Year()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create month")
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(nil, "Create month successfuly"))
}

func (h *MonthHandler) DeleteMonth(c *gin.Context) {
	idStr := c.Param("id")

	var id pgtype.UUID
	err := id.Scan(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid UUID"))
		return
	}

	err = h.service.DeleteMonth(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to Delete Month"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Month deleted"))

}

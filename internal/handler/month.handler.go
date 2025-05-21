package month_handler

import (
	"encoding/json"
	"net/http"
	"time"

	month_service "github.com/098765432m/monthly_planner_backend/internal/service/month"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
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

func (h *MonthHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateMonth)
	r.Delete("/{id}", h.DeleteMonth)

	return r
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

func (h *MonthHandler) CreateMonth(w http.ResponseWriter, r *http.Request) {
	var req CreateMonthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invadlid Month Date")
		return
	}

	parsedTime, err := time.Parse("01/2006", req.Month)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invadlid Month Date")
		return
	}

	err = h.service.CreateMonth(r.Context(), int8(parsedTime.Month()), int16(parsedTime.Year()))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create month")
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil, "Create month successfuly")

}

func (h *MonthHandler) DeleteMonth(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	var id pgtype.UUID
	err := id.Scan(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid UUID")
		return
	}

	err = h.service.DeleteMonth(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete month")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"Success": true,
		"Message": "Month deleted",
	})

}

package admin

import (
	"net/http"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler admin routes handler
type Handler struct {
	AdminUsecase usecase.AdminUsecase
}

// Login POST /api/admin/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

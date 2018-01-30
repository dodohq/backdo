package user

import (
	"net/http"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler handler for all user endpoints
type Handler struct {
	UserUsecase usecase.UserUsecase
}

// GetAllExistingAccounts GET /api/user
func (h *Handler) GetAllExistingAccounts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

// CreateNewAccount POST /api/user
func (h *Handler) CreateNewAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

// DeleteAccount DELETE /api/user/:id
func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

// Login POST /api/user/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

// ForgotPassword GET /api/user/forgot_password?email=
func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

// ResetPassword POST /api/user/reset_password
func (h *Handler) ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Implemented"))
}

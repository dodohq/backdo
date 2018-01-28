package admin

import (
	"encoding/json"
	"fmt"
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
	a := new(models.Admin)
	if err := helper.ReadRequestBody(r, a); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	token, err := h.AdminUsecase.Login(a.Email, a.Password)
	if err != (*models.HTTPError)(nil) {
		helper.RenderErr(w, err)
		return
	}

	helper.RenderJSON(w, map[string]string{"token": token})
}

// Dummy GET /api/admin
func (h *Handler) Dummy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	a := new(models.Admin)
	err := json.Unmarshal([]byte(r.Header.Get("User")), a)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(a)
	helper.RenderJSON(w, map[string]string{})
}

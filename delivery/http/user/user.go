package user

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

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
	accounts, httpErr := h.UserUsecase.GetAllUsers()
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string][]*models.User{"users": accounts})
}

// CreateNewAccount POST /api/user
func (h *Handler) CreateNewAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := helper.ReadRequestBody(r, u); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	emailReg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailReg.MatchString(u.Email) {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Email"))
		return
	} else if u.Password == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Password"))
		return
	} else if u.CompanyID < 1 {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Company ID"))
		return
	}

	httpErr := h.UserUsecase.CreateNewUser(u)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

// DeleteAccount DELETE /api/user/:id
func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	httpErr := h.UserUsecase.DeleteAnAccount(int64(ID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

// Login POST /api/user/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := helper.ReadRequestBody(r, u); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	token, httpErr := h.UserUsecase.Login(u.Email, u.Password)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{"token": token})
}

// ForgotPassword GET /api/user/forgot_password?email=
func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	emailQP := r.URL.Query()["email"]
	if len(emailQP) < 1 {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("No Email Provided"))
		return
	}
	email := emailQP[0]
	if email == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("No Email Provided"))
		return
	}

	httpErr := h.UserUsecase.ForgotPassword(email)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

// ResetPassword POST /api/user/reset_password
func (h *Handler) ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)

	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := helper.ReadRequestBody(r, u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	httpErr := h.UserUsecase.ResetPassword(u.Email, u.Password)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

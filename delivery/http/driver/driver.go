package driver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler driver endpoints handler
type Handler struct {
	DriverUsecase usecase.DriverUsecase
}

type verifyEndPointsBody struct {
	Via              string `json:"via" schema:"via"`
	CountryCode      int    `json:"country_code" schema:"country_code"`
	PhoneNumber      string `json:"phone_number" schema:"phone_number"`
	VerificationCode string `json:"verification_code" schema:"verification_code"`
}

// GetAllDriversOfCompany GET /api/driver
func (h *Handler) GetAllDriversOfCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	cID := u.CompanyID
	drivers, httpErr := h.DriverUsecase.GetAllDriversOfCompany(cID)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string][]*models.Driver{"drivers": drivers})
}

// CreateNewDriverProfile POST /api/driver
func (h *Handler) CreateNewDriverProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	d := new(models.Driver)
	if err := helper.ReadRequestBody(r, d); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	if d.Name == "" || d.PhoneNumber == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Information Provided"))
		return
	}
	d.CompanyID = u.CompanyID

	d, httpErr := h.DriverUsecase.CreateDriverProfile(d)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, d)
}

// UpdateDriverProfile PUT /api/driver
func (h *Handler) UpdateDriverProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.RenderErr(w, models.NewErrorInternalServer("Not Impelemented"))
}

// DeleteDriverProfile DELETE /api/driver/:id
func (h *Handler) DeleteDriverProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	dID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}
	d, httpErr := h.DriverUsecase.GetDriverByID(int64(dID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	} else if d.CompanyID != u.CompanyID {
		helper.RenderErr(w, models.NewErrorUnauthorized("No Access"))
		return
	}
	httpErr = h.DriverUsecase.RemoveDriverProfile(int64(dID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]interface{}{})
}

// RequestVerificationCode POST /api/driver/verification_code
// via is "sms" or "call"
func (h *Handler) RequestVerificationCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := new(verifyEndPointsBody)
	if err := helper.ReadRequestBody(r, body); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	if body.Via != "sms" && body.Via != "call" {
		body.Via = "sms"
	} else if body.CountryCode == 0 {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Country Code"))
		return
	} else if body.PhoneNumber == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Phone Number"))
		return
	}

	// check if the driver account exist
	_, httpErr := h.DriverUsecase.GetDriverByPhoneNumber(fmt.Sprintf("+%d%s", body.CountryCode, body.PhoneNumber))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	httpErr = h.DriverUsecase.RequestForVerificationCode(body.Via, body.CountryCode, body.PhoneNumber)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

// VerifyPhoneNumber POST /api/driver/check_verification
func (h *Handler) VerifyPhoneNumber(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := new(verifyEndPointsBody)
	if err := helper.ReadRequestBody(r, body); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	if body.CountryCode == 0 {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Country Code"))
		return
	} else if body.PhoneNumber == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Phone Number"))
		return
	} else if body.VerificationCode == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Invalid Verification Code"))
		return
	}

	_, httpErr := h.DriverUsecase.GetDriverByPhoneNumber(fmt.Sprintf("+%d%s", body.CountryCode, body.PhoneNumber))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	token, httpErr := h.DriverUsecase.AuthenticateDriver(body.CountryCode, body.PhoneNumber, body.VerificationCode)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{"token": token})
}

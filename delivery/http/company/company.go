package company

import (
	"net/http"
	"strconv"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler company routes hanlder
type Handler struct {
	CompanyUsecase usecase.CompanyUsecase
}

// GetAllCompanies GET /api/company
func (h *Handler) GetAllCompanies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allCompanies, httpErr := h.CompanyUsecase.GetAllCompanies()
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string][]*models.Company{"companies": allCompanies})
}

// OnboardNewCompany POST /api/company
func (h *Handler) OnboardNewCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := new(models.Company)
	if err := helper.ReadRequestBody(r, c); err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	if c.Name == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("No Name Provided"))
		return
	} else if c.ContactNumber == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("No Contact Number Provided"))
		return
	}

	newC, httpErr := h.CompanyUsecase.OnboardNewCompany(c)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, newC)
}

// DeleteACompany DELET /api/company/:id
func (h *Handler) DeleteACompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
		return
	}

	_, httpErr := h.CompanyUsecase.DeleteACompany(int64(cID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]string{})
}

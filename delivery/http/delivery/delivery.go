package delivery

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler handler for all delivery endpoints
type Handler struct {
	DeliveryUsecase usecase.DeliveryUsecase
}

// GetAllDeliveries GET /api/delivery
func (h *Handler) GetAllDeliveries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cID := u.CompanyID
	deliveries, httpErr := h.DeliveryUsecase.GetDeliveriesByCompanyID(cID)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string][]*models.Delivery{"deliveries": deliveries})
}

// CreateNewDelivery POST /api/delivery
func (h *Handler) CreateNewDelivery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	d := new(models.Delivery)
	err := helper.ReadRequestBody(r, d)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else if d.ContactNumber == "" || d.CustomerName == "" {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Not Enough Info Provided"))
		return
	}
	d.CompanyID = u.CompanyID

	d, httpErr := h.DeliveryUsecase.CreateNewDelivery(d)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]*models.Delivery{"delivery": d})
}

// BulkCreateDeliveries POST /api/delivery/bulk
func (h *Handler) BulkCreateDeliveries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	filePath, err := helper.ReadFileUpload(r, "csv")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fileReader, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer fileReader.Close()

	csvReader := csv.NewReader(fileReader)
	records, err := csvReader.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else if len(records) < 2 {
		helper.RenderErr(w, models.NewErrorUnprocessableEntity("Empty CSV File"))
		return
	}

	actualData := records[1:]
	deliveryList := make([]*models.Delivery, 0)
	for i, v := range actualData {
		if v[0] == "" || v[1] == "" {
			helper.RenderErr(w, models.NewErrorUnprocessableEntity(fmt.Sprintf("Not Enough Info In Row %d", i+1)))
			return
		}
		d := &models.Delivery{
			CustomerName:  v[0],
			ContactNumber: v[1],
			CompanyID:     u.CompanyID,
		}
		deliveryList = append(deliveryList, d)
	}

	deliveryList, httpErr := h.DeliveryUsecase.BulkCreateDeliveries(deliveryList)
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string][]*models.Delivery{"deliveries": deliveryList})
}

// DeleteADelivery DELETE /api/delivery/:id
func (h *Handler) DeleteADelivery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := new(models.User)
	if err := json.Unmarshal([]byte(r.Header.Get("User")), u); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	d, httpErr := h.DeliveryUsecase.GetDeliveryByID(int64(dID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	} else if d.CompanyID != u.CompanyID {
		helper.RenderErr(w, models.NewErrorUnauthorized("No Access"))
		return
	}
	httpErr = h.DeliveryUsecase.DeleteDelivery(int64(dID))
	if httpErr != (*models.HTTPError)(nil) {
		helper.RenderErr(w, httpErr)
		return
	}

	helper.RenderJSON(w, map[string]interface{}{})
}

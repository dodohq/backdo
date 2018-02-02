package delivery

import (
	"fmt"
	"sync"

	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
)

type privateDeliveryUsecase struct {
	deliveryRepo repository.DeliveryRepository
	companyRepo  repository.CompanyRepository
}

// GetAllDeliveries function to retrieve all deliveries
func (u *privateDeliveryUsecase) GetAllDeliveries() ([]*models.Delivery, *models.HTTPError) {
	deliveryList, httpErr := u.deliveryRepo.GetAllDeliveries()
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	}

	wg := new(sync.WaitGroup)
	for i, delivery := range deliveryList {
		wg.Add(1)
		go func(i int, delivery *models.Delivery) {
			c, httpErr := u.companyRepo.GetCompanyByID(delivery.CompanyID)
			if httpErr == (*models.HTTPError)(nil) {
				deliveryList[i].Company = c
			}
			wg.Done()
		}(i, delivery)
	}
	wg.Wait()

	return deliveryList, nil
}

// GetDeliveriesByCompanyID function to retrieve all deliveries belonging to a company
func (u *privateDeliveryUsecase) GetDeliveriesByCompanyID(id int64) ([]*models.Delivery, *models.HTTPError) {
	deliveryList, httpErr := u.deliveryRepo.GetDeliveriesByCompanyID(id)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	}

	c, httpErr := u.companyRepo.GetCompanyByID(id)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	}

	for i := range deliveryList {
		deliveryList[i].Company = c
	}

	return deliveryList, nil
}

// GetDeliveryByID function to retrieve a single delivery by its ID
func (u *privateDeliveryUsecase) GetDeliveryByID(id int64) (*models.Delivery, *models.HTTPError) {
	d, httpErr := u.deliveryRepo.GetDeliveryByID(id)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if d == (*models.Delivery)(nil) {
		return nil, models.NewErrorNotFound(fmt.Sprintf("Delivery with ID %d Doesnt Exist", id))
	}

	return d, nil
}

// CreateNewDelivery function to create a new delivery
func (u *privateDeliveryUsecase) CreateNewDelivery(d *models.Delivery) (*models.Delivery, *models.HTTPError) {
	return u.deliveryRepo.InsertDelivery(d)
}

// BulkCreateDeliveries function to create new deliveries in bulk
func (u *privateDeliveryUsecase) BulkCreateDeliveries(list []*models.Delivery) ([]*models.Delivery, *models.HTTPError) {
	return u.deliveryRepo.BulkInsertDelivery(list)
}

// DeleteDelivery function to delete a delivery by id
func (u *privateDeliveryUsecase) DeleteDelivery(id int64) *models.HTTPError {
	d, httpErr := u.deliveryRepo.GetDeliveryByID(id)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if d == (*models.Delivery)(nil) {
		return models.NewErrorNotFound(fmt.Sprintf("Delivery with ID %d Doesnt Exist", id))
	}

	return u.deliveryRepo.DeleteDelivery(id)
}

// NewDeliveryUsecase generate new delivery usecase
func NewDeliveryUsecase(deliveryRepo repository.DeliveryRepository, companyRepo repository.CompanyRepository) usecase.DeliveryUsecase {
	return &privateDeliveryUsecase{
		deliveryRepo,
		companyRepo,
	}
}

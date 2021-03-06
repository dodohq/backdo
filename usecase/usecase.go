package usecase

import "github.com/dodohq/backdo/models"

// AdminUsecase usecase interface for handling admin endpoints
type AdminUsecase interface {
	Login(email, password string) (string, *models.HTTPError)
}

// CompanyUsecase usecase interface for handling company endpoints
type CompanyUsecase interface {
	GetAllCompanies() ([]*models.Company, *models.HTTPError)
	OnboardNewCompany(c *models.Company) (*models.Company, *models.HTTPError)
	DeleteACompany(id int64) (bool, *models.HTTPError)
}

// UserUsecase usecase interface for handling user endpoints
type UserUsecase interface {
	GetAllUsers() ([]*models.User, *models.HTTPError)
	CreateNewUser(u *models.User) *models.HTTPError
	DeleteAnAccount(id int64) *models.HTTPError
	Login(email, password string) (string, *models.HTTPError)
	ForgotPassword(email string) *models.HTTPError
	ResetPassword(email, newPassword string) *models.HTTPError
}

// DeliveryUsecase usecase interface for handling delivery endpoints
type DeliveryUsecase interface {
	GetAllDeliveries() ([]*models.Delivery, *models.HTTPError)
	GetDeliveriesByCompanyID(id int64) ([]*models.Delivery, *models.HTTPError)
	GetDeliveryByID(id int64) (*models.Delivery, *models.HTTPError)
	CreateNewDelivery(d *models.Delivery) (*models.Delivery, *models.HTTPError)
	BulkCreateDeliveries(list []*models.Delivery) ([]*models.Delivery, *models.HTTPError)
	DeleteDelivery(id int64) *models.HTTPError
}

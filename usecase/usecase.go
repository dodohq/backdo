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

// UserUsecase usercase interface for handling user endpoints
type UserUsecase interface {
}

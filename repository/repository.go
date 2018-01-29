package repository

import "github.com/dodohq/backdo/models"

// AdminRepository interface for abstraction against third parties interaction
type AdminRepository interface {
	GetByEmail(email string) (*models.Admin, *models.HTTPError)
	GenerateJWT(a *models.Admin) (string, *models.HTTPError)
}

// CompanyRepository interface for abstraction against third parties interaction
type CompanyRepository interface {
	GetAllCompany() ([]*models.Company, *models.HTTPError)
	InsertNewCompany(c *models.Company) (*models.Company, *models.HTTPError)
	DeleteACompany(id int64) (bool, *models.HTTPError)
}

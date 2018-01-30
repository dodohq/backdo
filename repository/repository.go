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
	GetCompanyByID(id int64) (*models.Company, *models.HTTPError)
	InsertNewCompany(c *models.Company) (*models.Company, *models.HTTPError)
	DeleteACompany(id int64) (bool, *models.HTTPError)
}

// UserRepository interface for abstraction against third party interaction
type UserRepository interface {
	FetchAllUsers() ([]*models.User, *models.HTTPError)
	GetUserByEmail(email string) (*models.User, *models.HTTPError)
	UpdateUser(u *models.User) *models.HTTPError
	InsertNewUser(u *models.User) (*models.User, *models.HTTPError)
	DeleteUser(id int64) *models.HTTPError
	GenerateJWT(u *models.User) (string, *models.HTTPError)
	SendEmailToUser(u *models.User, subject, body string) *models.HTTPError
}

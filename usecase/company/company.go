package company

import (
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
)

type privateCompanyUsecase struct {
	companyRepo repository.CompanyRepository
}

// GetAllCompanies get all Do-Do's logistic partners
func (u *privateCompanyUsecase) GetAllCompanies() ([]*models.Company, *models.HTTPError) {
	return u.companyRepo.GetAllCompany()
}

// OnboardNewCompany as name suggested
func (u *privateCompanyUsecase) OnboardNewCompany(c *models.Company) (*models.Company, *models.HTTPError) {
	return u.companyRepo.InsertNewCompany(c)
}

// DeleteACompany when a company is no longer Do-Do partner
func (u *privateCompanyUsecase) DeleteACompany(id int64) (bool, *models.HTTPError) {
	return u.companyRepo.DeleteACompany(id)
}

// NewCompanyUsecase genereate new company usecase
func NewCompanyUsecase(companyRepo repository.CompanyRepository) usecase.CompanyUsecase {
	return &privateCompanyUsecase{
		companyRepo,
	}
}

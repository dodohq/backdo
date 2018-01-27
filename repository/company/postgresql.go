package company

import "github.com/dodohq/backdo/models"

// GetAllCompany get all partner companies
func (r *privateCompanyRepo) GetAllCompany() ([]*models.Company, *models.HTTPError) {
	return []*models.Company{}, models.NewErrorInternalServer("Not Implemented")
}

// InsertNewCompany onboard a new partner company
func (r *privateCompanyRepo) InsertNewCompany(*models.Company) (*models.Company, *models.HTTPError) {
	return nil, models.NewErrorInternalServer("Not Implemented")
}

// DeleteACompany delete a company from db
func (r *privateCompanyRepo) DeleteACompany(id int64) (bool, *models.HTTPError) {
	return false, models.NewErrorInternalServer("Not Implemented")
}

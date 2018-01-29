package usecase_test

import (
	"testing"

	"github.com/dodohq/backdo/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCompanies(t *testing.T) {
	expected := []*models.Company{
		&models.Company{ID: 1, Name: "Aramex", ContactNumber: "12345678"},
		&models.Company{ID: 2, Name: "DHL", ContactNumber: "87654321"},
	}

	companyRepoMock.On("GetAllCompany").Return(expected, (*models.HTTPError)(nil))
	actual, err := companyUsecase.GetAllCompanies()
	if err != (*models.HTTPError)(nil) {
		t.Errorf(err.Error())
		return
	}

	assert.Equal(t, actual, expected)
}

func TestOnboardNewCompany(t *testing.T) {
	newC := &models.Company{Name: "Aramex", ContactNumber: "12345678"}
	expected := &models.Company{ID: 1, Name: newC.Name, ContactNumber: newC.ContactNumber}
	companyRepoMock.On("InsertNewCompany", newC).Return(expected, (*models.HTTPError)(nil))

	actual, err := companyUsecase.OnboardNewCompany(newC)
	if err != (*models.HTTPError)(nil) {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, actual, expected)
}

func TestDeleteACompany(t *testing.T) {
	var ID int64 = 1
	companyRepoMock.On("DeleteACompany", ID).Return(true, (*models.HTTPError)(nil))

	actual, err := companyUsecase.DeleteACompany(ID)
	if err != (*models.HTTPError)(nil) {
		t.Errorf(err.Error())
		return
	}

	assert.True(t, actual)
}

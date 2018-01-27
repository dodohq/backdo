package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository/company"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllCompany(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	expected := []*models.Company{
		&models.Company{Name: "Aramex", ContactNumber: "12345678"},
		&models.Company{Name: "DHL", ContactNumber: "87654321"},
	}
	rows := sqlMock.NewRows(
		[]string{"name", "contact_numer"},
	).AddRow(
		expected[0].Name, expected[1].ContactNumber,
	).AddRow(
		expected[0].Name, expected[1].ContactNumber,
	)
	mock.ExpectQuery(`SELECT \* FROM companies`).WillReturnRows(rows)

	mockCompanyRepo := company.NewCompanyRepository(db)

	actual, err := mockCompanyRepo.GetAllCompany()
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	assert.Equal(t, expected, actual)
}

func TestInsertNewCompany(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	expected := &models.Company{ID: 1, Name: "Aramex", ContactNumber: "12345678"}
	result := sqlMock.NewResult(expected.ID, 1)
	mock.ExpectExec(`INSERT INTO companies`).WillReturnResult(result)

	mockCompanyRepo := company.NewCompanyRepository(db)

	actual, err := mockCompanyRepo.InsertNewCompany(expected)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	assert.Equal(t, expected, actual)
}

func TestDeleteACompany(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	result := sqlMock.NewResult(0, 1)
	mock.ExpectExec(`UPDATE companies`).WillReturnResult(result)

	mockCompanyRepo := company.NewCompanyRepository(db)

	outcome, err := mockCompanyRepo.DeleteACompany(1)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	assert.True(t, outcome)
}

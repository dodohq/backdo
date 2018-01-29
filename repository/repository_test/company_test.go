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
		&models.Company{ID: 1, Name: "Aramex", ContactNumber: "12345678"},
		&models.Company{ID: 2, Name: "DHL", ContactNumber: "87654321"},
	}
	rows := sqlMock.NewRows(
		[]string{"id", "name", "contact_numer"},
	).AddRow(
		expected[0].ID, expected[0].Name, expected[0].ContactNumber,
	).AddRow(
		expected[1].ID, expected[1].Name, expected[1].ContactNumber,
	)
	mock.ExpectQuery(`SELECT (.+) FROM companies`).WillReturnRows(rows)

	mockCompanyRepo := company.NewCompanyRepository(db)

	actual, err := mockCompanyRepo.GetAllCompany()
	if err != (*models.HTTPError)(nil) {
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
	result := sqlMock.NewRows([]string{"id"}).AddRow(expected.ID)
	mock.ExpectQuery(`INSERT INTO companies`).WillReturnRows(result)

	mockCompanyRepo := company.NewCompanyRepository(db)

	actual, err := mockCompanyRepo.InsertNewCompany(expected)
	if err != (*models.HTTPError)(nil) {
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
	mock.ExpectPrepare(`UPDATE companies SET deleted = TRUE WHERE id = \$1`)
	mock.ExpectExec(`UPDATE companies SET deleted = TRUE WHERE id = \$1`).WillReturnResult(result)

	mockCompanyRepo := company.NewCompanyRepository(db)

	outcome, err := mockCompanyRepo.DeleteACompany(1)
	if err != (*models.HTTPError)(nil) {
		t.Fatalf(err.Error())
		return
	}
	assert.True(t, outcome)
}

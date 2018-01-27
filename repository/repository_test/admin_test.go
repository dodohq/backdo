package repository_test

import (
	"testing"

	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository/admin"
	"github.com/stretchr/testify/assert"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByEmail(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error %s happened when opening a stub db connection", err)
	}
	defer db.Close()

	fakeEmail := "stanley@dodo.tech"
	fakePassword := "stan123"
	rows := sqlMock.NewRows([]string{"email", "password"}).AddRow(fakeEmail, fakePassword)
	mock.ExpectQuery(`SELECT \* FROM admins WHERE email = \$1`).WillReturnRows(rows)

	mockAdminRepo := admin.NewAdminRepository(db)

	actual, err := mockAdminRepo.GetByEmail("stanley@dodo.tech")
	if err != (*models.HTTPError)(nil) {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, &models.Admin{Email: fakeEmail, Password: fakePassword}, actual)
}

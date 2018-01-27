package usecase_test

import (
	"github.com/dodohq/backdo/usecase"
	"github.com/dodohq/backdo/usecase/admin"
	"github.com/dodohq/backdo/usecase/usecase_test/repomock"
)

var adminRepoMock *repomock.AdminRepoMock
var adminUsecase usecase.AdminUsecase

func init() {
	adminRepoMock = repomock.NewAdminRepoMock("password", "nonexistent@dodo.tech")
	adminUsecase = admin.NewAdminUsecase(adminRepoMock)
}

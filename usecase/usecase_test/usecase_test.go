package usecase_test

import (
	"github.com/dodohq/backdo/usecase"
	"github.com/dodohq/backdo/usecase/admin"
	"github.com/dodohq/backdo/usecase/company"
	"github.com/dodohq/backdo/usecase/usecase_test/repomock"
)

var adminRepoMock *repomock.AdminRepoMock
var adminUsecase usecase.AdminUsecase
var companyRepoMock *repomock.CompanyRepoMock
var companyUsecase usecase.CompanyUsecase

func init() {
	adminRepoMock = repomock.NewAdminRepoMock("password", "nonexistent@dodo.tech", "thisisafakejwt")
	adminUsecase = admin.NewAdminUsecase(adminRepoMock)
	companyRepoMock = repomock.NewCompanyRepoMock()
	companyUsecase = company.NewCompanyUsecase(companyRepoMock)
}

package company

import (
	"database/sql"

	"github.com/dodohq/backdo/repository"
)

type privateCompanyRepo struct {
	Conn *sql.DB
}

// NewCompanyRepository generate new company repo
func NewCompanyRepository(Conn *sql.DB) repository.CompanyRepository {
	return &privateCompanyRepo{
		Conn,
	}
}

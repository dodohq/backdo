package admin

import (
	"database/sql"

	"github.com/dodohq/backdo/repository"
)

type privateAdminRepo struct {
	Conn *sql.DB
}

// NewAdminRepository genereate new admin respository
func NewAdminRepository(Conn *sql.DB) repository.AdminRepository {
	return &privateAdminRepo{
		Conn,
	}
}

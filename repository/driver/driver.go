package driver

import (
	"database/sql"

	"github.com/dodohq/backdo/repository"
)

type privateDriverRepo struct {
	Conn *sql.DB
}

// NewDriverRepository generate new driver repository
func NewDriverRepository(Conn *sql.DB) repository.DriverRepository {
	return &privateDriverRepo{
		Conn,
	}
}

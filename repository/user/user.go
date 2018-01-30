package user

import (
	"database/sql"
	"net/smtp"

	"github.com/dodohq/backdo/repository"
)

type privateUserRepo struct {
	Conn *sql.DB
	Auth *smtp.Auth
}

// NewUserRepo generate new User Repository
func NewUserRepo(Conn *sql.DB, Auth *smtp.Auth) repository.UserRepository {
	return &privateUserRepo{
		Conn,
		Auth,
	}
}

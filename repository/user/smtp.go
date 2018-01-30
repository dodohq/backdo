package user

import (
	"net/smtp"
	"os"

	"github.com/dodohq/backdo/models"
)

// SendEmailToUser send an email to user
func (r *privateUserRepo) SendEmailToUser(u *models.User, body string) *models.HTTPError {
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		*r.Auth,
		os.Getenv("DODO_EMAIL"),
		[]string{u.Email},
		[]byte(body),
	)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}

	return nil
}

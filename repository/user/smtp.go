package user

import (
	"net/smtp"
	"os"

	"github.com/dodohq/backdo/models"
)

// SendEmailToUser send an email to user
func (r *privateUserRepo) SendEmailToUser(u *models.User, subject, body string) *models.HTTPError {
	msg := []byte(
		"From: " + os.Getenv("DODO_EMAIL") + "\n" +
			"To: " + u.Email + "\n" +
			"Subject: " + subject + "\n" +
			"\n" +
			body + "\n")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		*r.Auth,
		os.Getenv("DODO_EMAIL"),
		[]string{u.Email},
		msg,
	)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}

	return nil
}

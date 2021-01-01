package mailer

import (
	"context"

	"github.com/bensaufley/catalg/server/internal/log"
)

type EmailClient interface {
	PasswordReset(context.Context, PasswordResetEmailParams) error
}

type client struct {
	smtpServer string
	username   string
	password   string
}

func (c *client) Send(to string, from string, subject string, html string, text string) error {
	// TODO: set up send behavior
	log.WithFields(log.Fields{
		"to": to, "from": from, "subject": subject, "html": html, "text": text,
	}).Debug("client.Send called")
	return nil
}

var Client *client

func InitClient(smtpServer string, username string, password string) {
	Client = &client{smtpServer, username, password}
}

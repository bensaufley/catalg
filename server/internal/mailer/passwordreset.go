package mailer

import (
	"bytes"
	"context"
	"os"
	"path"
	"text/template"
)

var passwordResetTextTemplate = template.Must(template.ParseFiles(path.Join(os.Getenv("GOPATH"), "src/github.com/bensaufley/catalg/server/internal/mailer/templates/password-reset.txt.tmpl")))
var passwordResetHTMLTemplate = template.Must(template.ParseFiles(path.Join(os.Getenv("GOPATH"), "src/github.com/bensaufley/catalg/server/internal/mailer/templates/password-reset.html.tmpl")))

type PasswordResetEmailParams struct {
	Email    string
	Title    string
	Token    string
	Username string
}

func (c *client) PasswordReset(ctx context.Context, params PasswordResetEmailParams) error {
	if params.Title == "" {
		params.Title = "Catalg Password Reset"
	}

	var txt bytes.Buffer
	if err := passwordResetTextTemplate.Execute(&txt, params); err != nil {
		return err
	}

	var html bytes.Buffer
	if err := passwordResetHTMLTemplate.Execute(&html, params); err != nil {
		return err
	}

	// TODO: get a real "from" email
	return c.Send(params.Email, "catalg@example.com", params.Title, html.String(), txt.String())
}

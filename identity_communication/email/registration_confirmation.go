package email

import (
	"bytes"
	"html/template"
	"log/slog"
)

// RegistrationEmailConfirmationTemplate is the template for the registration confirmation email.
type RegistrationEmailConfirmationTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
}

// Content returns the content of the registration confirmation email
func (obj *RegistrationEmailConfirmationTemplate) Content() (html string, err error) {
	logger := slog.With("email_template", "registration confirmation template")
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/registration_confirmation.gohtml")
	if err != nil {
		logger.With("err", err).Error("could not parse template")
	}
	// run template engine
	err = t.Execute(&tpl, obj)
	if err != nil {
		logger.With("err", err).Error("could not execute template")
	}
	html = tpl.String()
	return

}

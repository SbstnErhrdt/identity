package email

import (
	"bytes"
	"html/template"
	"log/slog"
)

// RegistrationEmailTemplate is the template for the registration email
type RegistrationEmailTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
	HtmlTemplate   string
}

// DefaultRegistrationEmailResolver is the default resolver for the registration email
func DefaultRegistrationEmailResolver(origin, email, token string) RegistrationEmailTemplate {
	return RegistrationEmailTemplate{
		GlobalTemplate: DefaultGlobalTemplate,
		EmailOfNewUser: email,
		ContentText:    "Please click the button below to confirm your registration.",
		ButtonText:     "Confirm here by clicking",
		ButtonUrl:      "https://" + origin + "/identity/register/" + token,
	}
}

// Content returns the content of the registration email
func (obj *RegistrationEmailTemplate) Content() (html string, err error) {
	logger := slog.With("email_template", "registration template")
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/registration.gohtml")
	if err != nil {
		logger.With("err", err).Error("could not parse template")
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("registration-email").Parse(obj.HtmlTemplate)
		if errParse == nil {
			t = tNew
			logger.With("err", errParse).Error("could not parse provided html template")
		} else {
			logger.Warn("can not parse provided html template")
		}
	}
	// run template engine
	err = t.Execute(&tpl, obj)
	if err != nil {
		logger.With("err", err).Error("could not execute template")
	}
	html = tpl.String()
	return
}

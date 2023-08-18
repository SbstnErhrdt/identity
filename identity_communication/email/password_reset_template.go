package email

import (
	"bytes"
	"html/template"
	"log/slog"
)

// PasswordResetTemplate is the template for the password reset email.
type PasswordResetTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
	HtmlTemplate   string
}

// DefaultPasswordResetEmailResolver is the default resolver for the password reset email.
func DefaultPasswordResetEmailResolver(origin, email, token string) PasswordResetTemplate {
	return PasswordResetTemplate{
		GlobalTemplate: DefaultGlobalTemplate,
		EmailOfNewUser: email,
		ContentText:    "Please click the button below to reset your password.",
		ButtonText:     "Reset Password",
		ButtonUrl:      "https://" + origin + "/identity/password-reset/" + token,
	}
}

// Content returns the content of the email
func (obj *PasswordResetTemplate) Content() (html string, err error) {
	logger := slog.With("email_template", "password reset template")
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/reset_password.gohtml")
	if err != nil {
		logger.With("err", err).Error("could not parse template")
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("password-reset-email").Parse(obj.HtmlTemplate)
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

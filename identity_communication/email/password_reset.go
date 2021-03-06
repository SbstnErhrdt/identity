package email

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"html/template"
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
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/reset_password.gohtml")
	if err != nil {
		log.Error(err)
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("password-reset-email").Parse(obj.HtmlTemplate)
		if errParse == nil {
			t = tNew
		} else {
			log.Warn("can not parse provided html template")
		}
	}
	// run template engine
	err = t.Execute(&tpl, obj)
	if err != nil {
		log.Error(err)
	}
	html = tpl.String()
	return
}

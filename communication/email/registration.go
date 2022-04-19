package email

import (
	"bytes"
	"embed"
	log "github.com/sirupsen/logrus"
	"html/template"
)

//go:embed templates
var registrationTemplate embed.FS

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
		ButtonText:     "Confirm registration",
		ButtonUrl:      "https://" + origin + "/registration/confirm/" + token,
	}
}

// Content returns the content of the registration email
func (obj *RegistrationEmailTemplate) Content() (html string, err error) {
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(registrationTemplate, "templates/registration.gohtml")
	if err != nil {
		log.Error(err)
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("registration-email").Parse(obj.HtmlTemplate)
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

package email

import (
	"bytes"
	"embed"
	log "github.com/sirupsen/logrus"
	"html/template"
)

//go:embed templates
var registrationConfirmationTemplate embed.FS

type RegistrationEmailConfirmationTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
}

func (obj *RegistrationEmailConfirmationTemplate) Content() (html string, err error) {
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(registrationConfirmationTemplate, "templates/registration_confirmation.html")
	if err != nil {
		log.Error(err)
	}
	// run template engine
	err = t.Execute(&tpl, obj)
	if err != nil {
		log.Error(err)
	}
	html = tpl.String()
	return

}

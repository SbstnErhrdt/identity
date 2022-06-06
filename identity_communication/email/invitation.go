package email

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"html/template"
)

// InvitationEmailTemplate is the template for the invitation email
type InvitationEmailTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	FirstName      string
	LastName       string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
	HtmlTemplate   string
}

// DefaultInvitationEmailResolver is the default resolver for the invitation email
func DefaultInvitationEmailResolver(mandateUID uuid.UUID, clientUID *uuid.UUID, orgName, firstName, lastName, emailAddress, link string) InvitationEmailTemplate {
	return InvitationEmailTemplate{
		GlobalTemplate: DefaultGlobalTemplate,
		EmailOfNewUser: emailAddress,
		FirstName:      firstName,
		LastName:       lastName,
		ContentText:    fmt.Sprintf("Hello %s %s, you have been invited to join", firstName, lastName),
		ButtonText:     "Go to website",
		ButtonUrl:      link,
	}
}

// Content returns the content of the invitation email
func (obj *InvitationEmailTemplate) Content() (html string, err error) {
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/invitation.gohtml")
	if err != nil {
		log.Error(err)
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("invitation-email").Parse(obj.HtmlTemplate)
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

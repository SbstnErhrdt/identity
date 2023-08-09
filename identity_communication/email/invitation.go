package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
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
func DefaultInvitationEmailResolver(origin, firstName, lastName, emailAddress, link string) InvitationEmailTemplate {
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
	logger := slog.With("email_template", "invitation template")
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/invitation.gohtml")
	if err != nil {
		logger.With("err", err).Error("could not parse template")
		return
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("invitation-email").Parse(obj.HtmlTemplate)
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

package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
)

// CreationInvitationEmailTemplate is the template for the creation invitation email
type CreationInvitationEmailTemplate struct {
	GlobalTemplate
	EmailOfNewUser string
	FirstName      string
	LastName       string
	ContentText    string
	ButtonText     string
	ButtonUrl      string
	HtmlTemplate   string
}

// DefaultCreationInvitationEmailResolver is the default resolver for the invitation email
func DefaultCreationInvitationEmailResolver(origin, firstName, lastName, emailAddress, content, token string) CreationInvitationEmailTemplate {
	if len(content) == 0 {
		if len(firstName) == 0 && len(lastName) == 0 {
			content = fmt.Sprintf("Hello %s %s, you have been invited to join. Please click on the button to register.", firstName, lastName)
		} else {
			content = fmt.Sprintf("Hello, you have been invited to join. Please click on the button to register.")
		}
	}

	return CreationInvitationEmailTemplate{
		GlobalTemplate: DefaultGlobalTemplate,
		EmailOfNewUser: emailAddress,
		FirstName:      firstName,
		LastName:       lastName,
		ContentText:    content,
		ButtonText:     "Confirm Email and Login",
		ButtonUrl:      "https://" + origin + "/identity/invitation/" + token,
	}
}

// Content returns the content of the invitation email
func (obj *CreationInvitationEmailTemplate) Content() (html string, err error) {
	logger := slog.With("email_template", "invitation template")
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/creation-invitation.gohtml")
	if err != nil {
		logger.With("err", err).Error("could not parse template")
		return
	}
	// check if there are
	if len(obj.HtmlTemplate) > 0 {
		tNew, errParse := template.New("creation-invitation-email").Parse(obj.HtmlTemplate)
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

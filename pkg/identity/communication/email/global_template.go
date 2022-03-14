package email

import (
	"bytes"
	"embed"
	log "github.com/sirupsen/logrus"
	"html/template"
)

type GlobalTemplate struct {
	BackgroundColor     string
	PrimaryColor        string
	PrimaryBorderColor  string
	HeaderLogoUrl       string
	IntroText           string
	OutroText           string
	FooterCopyrightText string
	UnsubscribeLink     string
}

var DefaultGlobalTemplate = GlobalTemplate{
	BackgroundColor:     "#f6f6f6",
	PrimaryColor:        "#0099ff",
	PrimaryBorderColor:  "#0099ff",
	HeaderLogoUrl:       "",
	IntroText:           "",
	OutroText:           "",
	FooterCopyrightText: "",
	UnsubscribeLink:     "",
}

//go:embed templates
var cssTemplate embed.FS // css template

func (globalTemplate *GlobalTemplate) Style() (result template.HTML) {
	// init buffer
	var tpl bytes.Buffer
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(cssTemplate, "templates/style.gohtml")
	if err != nil {
		log.Error(err)
	}
	// run template engine
	err = t.Execute(&tpl, globalTemplate)
	if err != nil {
		log.Error(err)
	}
	result = template.HTML(tpl.String())
	return
}

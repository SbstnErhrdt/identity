package services

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/mail"
	"net/smtp"
	"os"
)

func SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) (err error) {
	// we used environment variables to load the
	// email address and the password from the shell
	// you can also directly assign the email address
	// and the password
	user := os.Getenv("SMPT_USER")
	password := os.Getenv("SMPT_PASSWORD")
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = senderAddress.String()
	headers["To"] = receiverAddress.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	// add mime
	message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// add content
	message += content

	// Connect to the SMTP Server
	servername := os.Getenv("SMPT_SERVER") + ":" + os.Getenv("SMPT_PORT")

	host, _, err := net.SplitHostPort(servername)
	if err != nil {
		log.Error(err)
		return err
	}

	auth := smtp.PlainAuth("", user, password, host)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require a ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil {
		log.Error(err)
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Error(err)
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Error(err)
		return err
	}

	// To && From
	if err = c.Mail(senderAddress.Address); err != nil {
		log.Error(err)
		return err
	}

	if err = c.Rcpt(receiverAddress.Address); err != nil {
		log.Error(err)
		return err
	}

	// TODO: BCC

	// Data
	w, err := c.Data()
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Error(err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	err = c.Quit()
	if err != nil {
		log.Error(err)
		return err
	}

	log.Println("email send successfully")
	return
}

package controllers

import (
	"encoding/csv"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// InviteCSV adds all users in a CSV file to the invite list
func InviteCSV(service IdentityService, mandateUID uuid.UUID, clientUID *uuid.UUID, filePath string, orgName, subject, link string) (err error) {
	logger := log.WithFields(
		log.Fields{
			"func":    "InviteCSV",
			"service": "identity",
			"file":    filePath,
		},
	)
	// Load a csv file.
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// Create a new reader.
	r := csv.NewReader(f)
	// Read in all the records.

	rowCounter := 0
	// iterate over the records

	for {
		record, errRead := r.Read()
		// Stop at EOF.
		if errRead == io.EOF {
			break
		}

		if errRead != nil {
			err = errRead
			return
		}

		// check if there are 3 columns
		if len(record) != 3 {
			err = errors.New("please use template to upload data")
			return
		}

		// skip the index
		if rowCounter == 0 {
			rowCounter++
			continue
		}

		firsName, lastName, emailAddress := record[0], record[1], record[2]
		// sanitize email
		emailAddress = SanitizeEmail(emailAddress)

		errInvite := InviteUser(service, mandateUID, clientUID, orgName, subject, firsName, lastName, emailAddress, link)
		if errInvite != nil {
			logger.WithFields(
				log.Fields{
					"firstName": firsName,
					"lastName":  lastName,
					"email":     emailAddress,
				},
			).WithError(errInvite).Info("could not invite")
		}
		rowCounter++
	}
	logger.WithField("amount", rowCounter).Info("email send")
	return
}

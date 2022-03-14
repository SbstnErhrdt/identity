package controllers

import "github.com/SbstnErhrdt/identity/pkg/identity/models"

// CreateTriplets takes an array of triplets and stores them in the database
func CreateTriplets(service IdentityService, triplets []*models.IdentityTriplet) (err error) {
	err = service.GetSQLClient().Create(triplets).Error
	return
}

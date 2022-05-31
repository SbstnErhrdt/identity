package identity_controllers

// CreateTriplets takes an array of triplets and stores them in the database
func CreateTriplets(service IdentityService, triplets []*identity_models.IdentityTriplet) (err error) {
	err = service.GetSQLClient().Create(triplets).Error
	return
}

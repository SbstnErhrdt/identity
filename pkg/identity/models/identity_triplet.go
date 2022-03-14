package models

type IdentityTriplet struct {
	// Metadata
	Base
	// Attributes
	Namespace string `gorm:"index" json:"namespace"`
	Object    string `gorm:"index" json:"object"`   // e.g. bookXYZ
	Relation  string `gorm:"index" json:"relation"` // e.g. can_write
	Subject   string `gorm:"index" json:"subject"`  // e.g. user
}

func (obj *IdentityTriplet) TableName() string {
	return "identity_triplets"
}

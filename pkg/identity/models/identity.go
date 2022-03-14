package models

type Identity struct {
	// Metadata
	Base
	// Attributes
	Salutation                 string `json:"salutation"`
	FirstName                  string `json:"firstName"`
	LastName                   string `json:"lastName"`
	Email                      string `json:"email"`
	BackupEmail                string `json:"backupEmail"`
	Phone                      string `json:"phone"`
	BackupPhone                string `json:"backupPhone"`
	Salt                       []byte `json:"-" xml:"-"`
	Password                   []byte `json:"-" xml:"-"`
	AcceptConditionsAndPrivacy bool   `json:"acceptConditionsAndPrivacy" xml:"-"`
	Active                     bool   `json:"active" xml:"-"`
	Cleared                    bool   `gorm:"default:0" json:"cleared" xml:"-"`
	Blocked                    bool   `gorm:"default:0" json:"blocked" xml:"-"`
}

func (obj *Identity) TableName() string {
	return "identity_identities"
}

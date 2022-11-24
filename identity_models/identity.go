package identity_models

import (
	"errors"
	"github.com/SbstnErhrdt/identity/security"
	"reflect"
)

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

// ErrNoPepper is returned if no pepper is provided
var ErrNoPepper = errors.New("no pepper")

// SetNewPassword sets a new password for the identity
func (obj *Identity) SetNewPassword(pepper, password string) (err error) {
	// Check for pepper
	if pepper == "" {
		err = ErrNoPepper
		return
	}
	// Hash pw and salt
	pw, salt := security.HashPassword(pepper, password, []byte{})
	obj.Password = pw
	obj.Salt = salt
	return
}

// CheckPassword checks the password for the identity
func (obj *Identity) CheckPassword(pepper, password string) bool {
	checkPassword, _ := security.HashPassword(pepper, password, obj.Salt)
	return reflect.DeepEqual(obj.Password, checkPassword)
}

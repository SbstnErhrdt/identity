package identity_models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID         int64      `json:"-" gorm:"primary_key;autoIncrement:true"`
	UID        uuid.UUID  `json:"UID" gorm:"type:varchar(36)" sql:"index"`
	CreatedAt  time.Time  `json:"created_at" sql:"index"`
	UpdatedAt  time.Time  `json:"update_at"`
	DeletedAt  *time.Time `json:"deleted_at" sql:"index"`
	ArchivedAt *time.Time `json:"archived_at" sql:"index"`
}

// BeforeCreate intercepts before the in the database
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now().UTC()
	base.UpdatedAt = time.Now().UTC()
	base.ID = 0
	// init uuid if its created
	if base.UID == uuid.Nil {
		base.UID = uuid.New()
	}
	return nil
}

// BeforeSave intercepts before the save in the database
func (base *Base) BeforeSave(tx *gorm.DB) error {
	base.UpdatedAt = time.Now().UTC()
	return nil
}

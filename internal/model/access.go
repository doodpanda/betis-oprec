package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccessPermission represents the access record of a wizard to a magic book.
type AccessPermission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	WitchID     uuid.UUID `gorm:"type:uuid;not null"`
	MagicBookID uuid.UUID `gorm:"type:uuid;not null"`
	AccessDate  time.Time `gorm:"not null"`

	// Relationships
	Witch     Witch     `gorm:"foreignKey:WitchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MagicBook MagicBook `gorm:"foreignKey:MagicBookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// BeforeCreate generates a UUID before creating a new record.
func (ap *AccessPermission) BeforeCreate(tx *gorm.DB) (err error) {
	ap.ID = uuid.New()
	return
}

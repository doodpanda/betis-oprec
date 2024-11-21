package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccessPermission represents the access record of a wizard to a magic book.
// AccessPermission represents the permissions granted to a witch for accessing a magic book.
//
// Fields:
// - ID: Unique identifier for the access permission (UUID). This is the primary key.
// - WitchID: Unique identifier for the witch (UUID). This field is not nullable and is indexed with MagicBookID to ensure uniqueness. It has foreign key constraints with cascading updates and deletions.
// - MagicBookID: Unique identifier for the magic book (UUID). This field is not nullable and is indexed with WitchID to ensure uniqueness. It has foreign key constraints with cascading updates and deletions.
// - PermitDate: The date when the permission was granted. This field is not nullable.

type AccessPermission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	WitchID     uuid.UUID `gorm:"type:uuid;not null;index:idx_witch_magicbook,unique;foreignKey:WitchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MagicBookID uuid.UUID `gorm:"type:uuid;not null;index:idx_witch_magicbook,unique;foreignKey:MagicBookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PermitDate  time.Time `gorm:"not null"`
}

// BeforeCreate generates a UUID before creating a new record.
func (ap *AccessPermission) BeforeCreate(tx *gorm.DB) (err error) {
	ap.ID = uuid.New()
	return
}

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MagicBook struct represents the structure of a magical book in the library.
// It contains information about the book's ID, title, type of magic, creation date, and availability status.
type MagicBook struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	MagicType   string    `gorm:"type:magic_type;not null"`
	CreatedDate time.Time `gorm:"type:timestamptz"`
	IsAvailable bool      `gorm:"default:true"`
}

func (b *MagicBook) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New() // Generate UUID using Go
	return
}

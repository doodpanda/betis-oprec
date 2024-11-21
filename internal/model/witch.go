package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Wizard represents the struct of a wizard in the system.
type Witch struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"type:varchar(255);not null"`
	Age  int       `gorm:"not null"`
	Rank string    `gorm:"type:magic_rank;not null"`
}

// BeforeCreate generates a UUID before creating a new record.
func (w *Witch) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}

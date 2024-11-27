package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        *uuid.UUID `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;type:timestamptz"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime;type:timestamptz"`
	DeletedAt time.Time  `json:"-" gorm:"index"`
}

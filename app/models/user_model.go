package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Email     string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PassHash  string    `db:"pass_hash" json:"passHash,omitempty" validate:"required,lte=255"`
}

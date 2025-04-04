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

type UserProfile struct {
	UserId   uuid.UUID `db:"user_id" json:"userId" validate:"required,uuid"`
	Nickname string    `db:"nickname" json:"nickname" validate:"required,lte=32"`
	Bio      string    `db:"bio" json:"bio" validate:"lte=620"`
}

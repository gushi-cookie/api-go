package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" validate:"required"`
	UpdatedAt time.Time `db:"updated_at"`
	Email     string    `db:"email" validate:"required,email,lte=255"`
	PassHash  string    `db:"pass_hash" validate:"required,lte=255"`
}

type UserProfile struct {
	UserId   uuid.UUID `db:"user_id" validate:"required,uuid"`
	Nickname string    `db:"nickname" validate:"required,lte=32"`
	Bio      string    `db:"bio" validate:"lte=620"`
}

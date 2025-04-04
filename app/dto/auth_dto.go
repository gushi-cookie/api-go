package dto

import "github.com/google/uuid"

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	Nickname string `json:"nickname" validate:"required,lte=32"`
	Bio      string `json:"bio" validate:"lte=620"`
}

type SignUpResponse struct {
	Message string  `json:"message"`
	User    NewUser `json:"user"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type SignInResponse struct {
	Message string `json:"message"`
	Tokens  Tokens `json:"tokens"`
}

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type NewUser struct {
	ID       uuid.UUID `json:"id"`
	Nickname string    `json:"nickname"`
	Bio      string    `json:"bio"`
}

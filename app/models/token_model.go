package models

type RenewTokens struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

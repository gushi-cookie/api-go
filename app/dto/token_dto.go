package dto

type RenewTokensRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RenewTokensResponse struct {
	Message string `json:"message"`
	Tokens  Tokens `json:"tokens"`
}

package dto

type LoginRequest struct {
	Nim      string `json:"nim" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

type CurrentUserResponse struct {
	Nim string `json:"nim"`
}

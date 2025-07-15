package domain_response

type AuthResponse struct {
	Token string `json:"token"`
}

type ValidateResetPasswordCodeResponse struct {
	IsValid bool `json:"is_valid"`
}

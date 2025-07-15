package dto

type AuthInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleAuthInputDto struct {
	Token string `json:"token"`
}

type ResetPasswordInputDto struct {
	Code        int    `json:"code"`
	NewPassword string `json:"new_password"`
	Email       string `json:"email"`
}

type GenerateResetPasswordCodeInputDto struct {
	Email string `json:"email"`
}

type ValidateResetPasswordCodeInputDto struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

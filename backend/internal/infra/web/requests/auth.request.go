package requests

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleAuthRequest struct {
	Token string `json:"token"`
}

type GenerateResetPasswordCodeRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Code        int    `json:"code"`
}

type ValidateResetPasswordCodeRequest struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

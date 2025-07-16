package dto

import "mime/multipart"

type PhotoUpload struct {
	File        multipart.File
	FileName    string
	FileSize    int64
	ContentType string
}
type CreateUserInputDto struct {
	Email    string
	Name     string
	Password *string
	Origin   string // e.g., "local" or "google"
}

type GetUserInputDto struct {
	Uuid string `json:"uuid"`
}

type UpdateUserInputDto struct {
	Uuid  string  `json:"uuid"`
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

type UserResetPasswordInputDto struct {
	Uuid        string `json:"uuid"`
	NewPassword []byte `json:"new_password"`
}

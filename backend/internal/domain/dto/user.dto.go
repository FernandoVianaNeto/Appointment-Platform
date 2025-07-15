package dto

import "mime/multipart"

type PhotoUpload struct {
	File        multipart.File
	FileName    string
	FileSize    int64
	ContentType string
}
type CreateUserInputDto struct {
	Email           string
	BirthDate       string
	Name            string
	Password        *string
	Origin          string // e.g., "local" or "google"
	ShippingAddress string
	BillingAddress  string
}

type GetUserInputDto struct {
	Uuid string `json:"uuid"`
}

type UpdateUserInputDto struct {
	Uuid            string  `json:"uuid"`
	Email           *string `json:"email"`
	BirthDate       *string `json:"birth_date"`
	Name            *string `json:"name"`
	ShippingAddress *string `json:"shipping_address"`
	BillingAddress  *string `json:"billing_address"`
}

type UserResetPasswordInputDto struct {
	Uuid        string `json:"uuid"`
	NewPassword []byte `json:"new_password"`
}

package requests

type CreateGoogleUserRequest struct {
	Email           string `json:"email"`
	BirthDate       string `json:"birth_date"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	GoogleSub       string `json:"google_sub"`
	ShippingAddress string `json:"shipping_address"`
	BillingAddress  string `json:"billing_address"`
}

type GetByUuidRequest struct {
	Uuid string `uri:"uuid"`
}

package domain_response

type GetUserProfileResponse struct {
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	BirthDate       string `json:"birth_date"`
	ShippingAddress string `json:"shipping_address"`
	BillingAddress  string `json:"billing_address"`
}

package entity

type User struct {
	Uuid            string  `json:"uuid"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	BirthDate       string  `json:"birth_date"`
	BillingAddress  string  `json:"billing_address"`
	ShippingAddress string  `json:"shipping_address"`
	Password        *[]byte `json:"password,omitempty"`
	AuthProvider    string  `json:"auth_provider"`
	GoogleSub       *string `json:"google_sub,omitempty"`
}

func NewUser(
	uuid string,
	email string,
	birthDate string,
	name string,
	password *[]byte,
	authProvider string,
	googleSub *string,
	shippingAddress string,
	billingAddress string,
) *User {
	entity := &User{
		Uuid:            uuid,
		Email:           email,
		Name:            name,
		Password:        password,
		BirthDate:       birthDate,
		AuthProvider:    authProvider,
		GoogleSub:       googleSub,
		ShippingAddress: shippingAddress,
		BillingAddress:  billingAddress,
	}
	return entity
}

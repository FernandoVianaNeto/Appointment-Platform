package mongo_repository

type UserModel struct {
	Uuid            string  `bson:"uuid" json:"uuid"`
	Email           string  `bson:"email" json:"email"`
	Name            string  `bson:"name" json:"name"`
	Password        *string `bson:"password,omitempty" json:"password,omitempty"`
	ShippingAddress string  `bson:"shipping_address" json:"shipping_address"`
	BillingAddress  string  `bson:"billing_address" json:"billing_address"`
	BirthDate       string  `bson:"birth_date" json:"birth_date"`
	AuthProvider    string  `bson:"auth_provider" json:"auth_provider"`
	GoogleSub       *string `bson:"google_sub,omitempty" json:"google_sub,omitempty"`
}

package mongo_repository

type UserModel struct {
	Uuid         string  `bson:"uuid" json:"uuid"`
	Email        string  `bson:"email" json:"email"`
	Name         string  `bson:"name" json:"name"`
	Password     *string `bson:"password,omitempty" json:"password,omitempty"`
	AuthProvider string  `bson:"auth_provider" json:"auth_provider"`
	GoogleSub    *string `bson:"google_sub,omitempty" json:"google_sub,omitempty"`
}

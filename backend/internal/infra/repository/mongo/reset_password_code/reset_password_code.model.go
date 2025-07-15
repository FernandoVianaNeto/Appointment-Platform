package reset_password_code_mongo_repository

type ResetPasswordCodeModel struct {
	UserUuid         string `bson:"user_uuid" json:"user_uuid"`
	Code             int    `bson:"code" json:"code"`
	CodeExpiration   string `bson:"code_expiration" json:"code_expiration"`
	AlreadyActivated bool   `bson:"already_activated" json:"already_activated"`
	Email            string `bson:"email" json:"email"`
	CreatedAt        string `bson:"created_at" json:"created_at"`
	ActivatedAt      string `bson:"activated_at" json:"activated_at"`
}

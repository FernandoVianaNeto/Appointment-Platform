package patient_mongo_repository

type PatientModel struct {
	Uuid      string  `json:"uuid" bson:"uuid"`
	UserUuid  string  `json:"user_uuid" bson:"user_uuid"`
	Name      string  `json:"name" bson:"name"`
	Phone     string  `json:"phone" bson:"phone"`
	Email     string  `json:"email" bson:"email"`
	Insurance string  `json:"insurance" bson:"insurance"`
	Address   *string `json:"address,omitempty" bson:"address"`
}

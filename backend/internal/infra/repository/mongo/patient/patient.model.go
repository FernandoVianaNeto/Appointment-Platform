package patient_mongo_repository

type PatientModel struct {
	Uuid      string  `json:"uuid" bson:"uuid"`
	Name      string  `json:"name" bson:"name"`
	Phone     string  `json:"phone" bson:"phone"`
	Insurance *string `json:"insurance" bson:"insurance"`
	Address   *string `json:"address,omitempty" bson:"address"`
	Email     *string `json:"email,omitempty" bson:"email"`
}

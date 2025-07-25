package entity

type Patient struct {
	Uuid      string  `json:"uuid"`
	UserUuid  string  `json:"user_uuid"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Insurance string  `json:"insurance"`
	Email     string  `json:"email"`
	Address   *string `json:"address,omitempty"`
}

func NewPatient(
	uuid string,
	user_uuid string,
	name string,
	phone string,
	insurance string,
	address *string,
	email string,
) *Patient {
	entity := &Patient{
		Uuid:      uuid,
		UserUuid:  user_uuid,
		Name:      name,
		Insurance: insurance,
		Address:   address,
		Email:     email,
		Phone:     phone,
	}
	return entity
}

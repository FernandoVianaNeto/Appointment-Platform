package entity

type User struct {
	Uuid         string  `json:"uuid"`
	Email        string  `json:"email"`
	Name         string  `json:"name"`
	Password     *[]byte `json:"password,omitempty"`
	AuthProvider string  `json:"auth_provider"`
	GoogleSub    *string `json:"google_sub,omitempty"`
}

func NewUser(
	uuid string,
	email string,
	name string,
	password *[]byte,
	authProvider string,
	googleSub *string,
) *User {
	entity := &User{
		Uuid:         uuid,
		Email:        email,
		Name:         name,
		Password:     password,
		AuthProvider: authProvider,
		GoogleSub:    googleSub,
	}
	return entity
}

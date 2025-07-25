package dto

type CreatePatientInputDto struct {
	Uuid      string  `json:"uuid"`
	UserUuid  string  `json:"user_uuid"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Insurance string  `json:"insurance"`
	Address   *string `json:"address"`
}

type ListPatientInputDto struct {
	Page        int     `json:"page"`
	UserUuid    string  `json:"user_uuid"`
	Name        *string `json:"name"`
	Uuid        *string `json:"uuid"`
	SearchInput *string `json:"search_input"`
	FilterType  *string `json:"filter_type"`
}

type EditPatientInputDto struct {
	Uuid        string  `json:"uuid"`
	PatientUuid string  `json:"patient_uuid"`
	Name        *string `json:"name"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Address     *string `json:"address"`
}

type DeletePatientInputDto struct {
	Uuids []string `json:"uuids"`
}

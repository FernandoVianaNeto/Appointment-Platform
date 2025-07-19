package dto

type CreatePatientInputDto struct {
	Uuid      string  `json:"uuid"`
	UserUuid  string  `json:"user_uuid"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Email     *string `json:"email"`
	Address   *string `json:"address"`
	Insurance *string `json:"insurance"`
}

type ListPatientInputDto struct {
	Page        int     `json:"page"`
	SearchInput *string `json:"search_input"`
	FilterType  *string `json:"filter_type"`
}

type EditPatientInputDto struct {
	Uuid    string  `json:"uuid"`
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Email   *string `json:"email"`
	Address *string `json:"address"`
}

type DeletePatientInputDto struct {
	Uuid *string `json:"uuid"`
}

package domain_response

type PatientData struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Insurance string `json:"insurance"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}

type ListPatientsResponse struct {
	Data     []PatientData      `json:"data"`
	Metadata PaginationMetadata `json:"metadata"`
}

package domain_response

type PatientData struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Insurance string `json:"insurance"`
	Phone     string `json:"phone"`
}

type ListPatientsResponse struct {
	Data     []PatientData      `json:"data"`
	Metadata PaginationMetadata `json:"metadata"`
}

package domain_response

type Patient struct {
	Name      string `json:"name"`
	Insurance string `json:"insurance"`
	Phone     string `json:"phone"`
}

type AppointmentData struct {
	Uuid      string  `json:"uuid"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Patient   Patient `json:"patient"`
	Procedure string  `json:"procedure"`
	Location  string  `json:"location"`
	Status    string  `json:"status"`
}

type ListAppointmentsResponse struct {
	Data     []AppointmentData  `json:"data"`
	Metadata PaginationMetadata `json:"metadata"`
}

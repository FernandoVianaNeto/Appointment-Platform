package requests

type ListAppointmentRequest struct {
	Page       string `form:"page"`
	SearchTerm string `form:"searchTerm"`
	FilterType string `form:"filterType"`
	Date       string `form:"date"`
}

type DeleteAppointmentRequest struct {
	Uuids []string `json:"uuids"`
}

type SetAppointmentStatusRequest struct {
	Status string `form:"status" required:"true"`
	Uuid   string `form:"uuid" required:"true"`
}

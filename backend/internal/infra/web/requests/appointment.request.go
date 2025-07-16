package requests

type ListAppointmentRequest struct {
	Page       string `form:"page"`
	SearchTerm string `form:"searchTerm"`
	FilterType string `form:"filterType"`
}

type DeleteAppointmentRequest struct {
	Uuid string `uri:"uuid"`
}

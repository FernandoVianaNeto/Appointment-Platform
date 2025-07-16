package requests

type ListPatientRequest struct {
	Page       string `form:"page"`
	SearchTerm string `form:"searchTerm"`
	FilterType string `form:"filterType"`
}

type DeletePatientRequest struct {
	Uuid string `uri:"uuid"`
}

package requests

type ListPatientRequest struct {
	Uuid       *string `form:"uuid"`
	Name       *string `form:"name"`
	Page       string  `form:"page"`
	SearchTerm *string `form:"searchTerm"`
	FilterType *string `form:"filterType"`
}

type DeletePatientRequest struct {
	Uuid string `uri:"uuid"`
}

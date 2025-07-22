package dto

type CreateAppointmentInputDto struct {
	Uuid        string `json:"uuid"`
	UserUuid    string `json:"user_uuid"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	PatientUuid string `json:"patient_uuid"`
	Insurance   string `json:"insurance"`
	Technician  string `json:"technician"`
	Location    string `json:"location"`
	Procedure   string `json:"procedure"`
}

type ListAppointmentInputDto struct {
	Page        int     `json:"page"`
	UserUuid    string  `json:"user_uuid"`
	Date        *string `json:"date"`
	SearchInput *string `json:"search_input"`
	FilterType  *string `json:"filter_type"`
}

type EditAppointmentInputDto struct {
	Uuid      string  `json:"uuid"`
	StartDate *string `json:"start_date"`
	EndDate   *string `json:"end_date"`
	Procedure *string `json:"procedure"`
	Status    *string `json:"status"`
}

type DeleteAppointmentInputDto struct {
	Uuids []string `json:"uuids"`
}

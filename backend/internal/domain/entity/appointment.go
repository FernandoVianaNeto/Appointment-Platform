package entity

type Appointment struct {
	Uuid        string `json:"uuid"`
	UserUuid    string `json:"user_uuid"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	PatientUuid string `json:"patient_uuid"`
	PatientName string `json:"patient_name"`
	Insurance   string `json:"insurance"`
	Technician  string `json:"technician"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Procedure   string `json:"procedure"`
}

func NewAppointment(
	uuid string,
	userUuid string,
	startDate string,
	endDate string,
	patientUuid string,
	patientName string,
	insurance string,
	technician string,
	location string,
	procedure string,
) *Appointment {
	entity := &Appointment{
		Uuid:        uuid,
		UserUuid:    userUuid,
		StartDate:   startDate,
		EndDate:     endDate,
		PatientUuid: patientUuid,
		PatientName: patientName,
		Insurance:   insurance,
		Technician:  technician,
		Location:    location,
		Status:      "not_confirmed",
		Procedure:   procedure,
	}
	return entity
}

package appointment_mongo_repository

type AppointmentModel struct {
	Uuid         string `json:"uuid" bson:"uuid"`
	UserUuid     string `json:"user_uuid" bson:"user_uuid"`
	StartDate    string `json:"start_date" bson:"start_date"`
	EndDate      string `json:"end_date" bson:"end_date"`
	PatientUuid  string `json:"patient_uuid" bson:"patient_uuid"`
	Insurance    string `json:"insurance" bson:"insurance"`
	Technician   string `json:"technician" bson:"technician"`
	Location     string `json:"location" bson:"location"`
	Status       string `json:"status" bson:"status"`
	Procedure    string `json:"procedure" bson:"procedure"`
	ReminderSent bool   `json:"reminder_sent" bson:"reminder_sent"`
}

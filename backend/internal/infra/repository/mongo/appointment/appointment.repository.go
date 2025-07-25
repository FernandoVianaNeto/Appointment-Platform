package appointment_mongo_repository

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppointmentRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewAppointmentRepository(db *mongo.Database) domain_repository.AppointmentRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.AppointmentCollection)

	return &AppointmentRepository{
		db:         db,
		collection: collection,
	}
}

func (f *AppointmentRepository) Create(ctx context.Context, input entity.Appointment) error {
	_, err := f.GetByTimeAndTechnician(ctx, input.StartDate, input.EndDate, input.Technician)

	if err != mongo.ErrNoDocuments {
		return errors.New("could not create a new appointment")
	}

	appointmentEntity := AppointmentModel{
		UserUuid:     input.UserUuid,
		Uuid:         input.Uuid,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		PatientUuid:  input.PatientUuid,
		Insurance:    input.Insurance,
		Technician:   input.Technician,
		Location:     input.Location,
		Status:       input.Status,
		Procedure:    input.Procedure,
		ReminderSent: input.ReminderSent,
	}

	_, err = f.collection.InsertOne(ctx, appointmentEntity)

	if err != nil {
		return err
	}

	return nil
}

func (f *AppointmentRepository) GetByTimeAndTechnician(ctx context.Context, startDate string, endDate string, technician string) (*entity.Appointment, error) {
	var model AppointmentModel

	filter := bson.M{
		"startDate": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
		"technician": technician,
	}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		return nil, err
	}

	response := entity.Appointment{
		Uuid:        model.Uuid,
		UserUuid:    model.UserUuid,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		PatientUuid: model.PatientUuid,
		Insurance:   model.Insurance,
		Technician:  model.Technician,
		Location:    model.Location,
		Status:      model.Status,
		Procedure:   model.Procedure,
	}

	return &response, err
}

func (f *AppointmentRepository) List(ctx context.Context, input dto.ListAppointmentInputDto) ([]entity.Appointment, error) {
	filters := buildListFilters(input)

	limit := int64(domain_response.DEFAULT_ITEMS_PER_PAGE)
	skip := int64((input.Page - 1)) * limit

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetSort(bson.D{
		{"start_date", 1},
		{"technician", 1},
	})

	cursor, err := f.collection.Find(ctx, filters, opts)
	if err != nil {
		return []entity.Appointment{}, err
	}
	defer cursor.Close(ctx)

	var appointments []AppointmentModel
	if err := cursor.All(ctx, &appointments); err != nil {
		return []entity.Appointment{}, err
	}

	entitiesAppointment := make([]entity.Appointment, 0, len(appointments))
	for _, appointment := range appointments {
		entitiesAppointment = append(entitiesAppointment, entity.Appointment{
			Uuid:        appointment.Uuid,
			UserUuid:    appointment.UserUuid,
			StartDate:   appointment.StartDate,
			EndDate:     appointment.EndDate,
			PatientUuid: appointment.PatientUuid,
			Insurance:   appointment.Insurance,
			Technician:  appointment.Technician,
			Location:    appointment.Location,
			Status:      appointment.Status,
			Procedure:   appointment.Procedure,
		})
	}

	return entitiesAppointment, nil
}

func (f *AppointmentRepository) Edit(ctx context.Context, input dto.EditAppointmentInputDto) error {
	updateFields := bson.M{}

	if input.StartDate != nil {
		updateFields["start_date"] = input.StartDate
	}
	if input.EndDate != nil {
		updateFields["end_date"] = input.EndDate
	}
	if input.Procedure != nil {
		updateFields["procedure"] = input.Procedure
	}
	if input.Status != nil {
		updateFields["status"] = input.Status
	}

	if len(updateFields) == 0 {
		return nil
	}

	update := bson.M{"$set": updateFields}

	filter := bson.M{"uuid": input.Uuid}

	err := f.collection.FindOneAndUpdate(ctx, filter, update).Err()
	return err
}

func (f *AppointmentRepository) Delete(ctx context.Context, uuid string) {
	filter := bson.M{"uuid": uuid}

	f.collection.FindOneAndDelete(ctx, filter)
}

func (f *AppointmentRepository) DeleteMany(ctx context.Context, ids []string) error {
	filter := bson.M{
		"uuid": bson.M{
			"$in": ids,
		},
	}

	_, err := f.collection.DeleteMany(ctx, filter)
	return err
}

func (f *AppointmentRepository) CountDocuments(ctx context.Context, input dto.ListAppointmentInputDto) (int64, error) {
	filters := buildListFilters(input)

	total, err := f.collection.CountDocuments(ctx, filters)

	return total, err
}

func (f *AppointmentRepository) GetNextAppointments(ctx context.Context, window time.Duration) (*[]entity.Appointment, error) {
	now := time.Now().UTC()
	from := time.Now().Format("2006-01-02T15:04")
	to := time.Now().Add(24 * time.Hour).Format("2006-01-02T15:04")

	fmt.Println(now, from, to)

	filter := bson.M{
		"start_date": bson.M{
			"$gte": from,
			"$lt":  to,
		},
		"reminder_sent": false,
	}

	cursor, err := f.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var appointments []AppointmentModel
	if err := cursor.All(ctx, &appointments); err != nil {
		return nil, err
	}

	entitiesAppointment := make([]entity.Appointment, 0, len(appointments))
	for _, appointment := range appointments {
		entitiesAppointment = append(entitiesAppointment, entity.Appointment{
			Uuid:        appointment.Uuid,
			UserUuid:    appointment.UserUuid,
			StartDate:   appointment.StartDate,
			EndDate:     appointment.EndDate,
			PatientUuid: appointment.PatientUuid,
			Insurance:   appointment.Insurance,
			Technician:  appointment.Technician,
			Location:    appointment.Location,
			Status:      appointment.Status,
			Procedure:   appointment.Procedure,
		})
	}

	return &entitiesAppointment, err
}

func (f *AppointmentRepository) UpdateReminderSent(ctx context.Context, uuid string) error {
	filter := bson.M{
		"uuid":          uuid,
		"reminder_sent": false,
	}

	update := bson.M{
		"$set": bson.M{
			"reminder_sent": true,
		},
	}

	err := f.collection.FindOneAndUpdate(ctx, filter, update).Err()
	return err
}

func (f *AppointmentRepository) UpdateStatus(ctx context.Context, status string, uuid string) error {
	filter := bson.M{
		"uuid": uuid,
	}

	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	err := f.collection.FindOneAndUpdate(ctx, filter, update).Err()
	return err
}

func buildListFilters(input dto.ListAppointmentInputDto) bson.M {
	filters := bson.M{
		"user_uuid": input.UserUuid,
	}

	var start string
	if input.Date == nil {
		now := time.Now()
		startUTC := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		start = startUTC.Format("2006-01-02")

	} else {
		start = *input.Date
	}

	filters["start_date"] = bson.M{
		"$regex":   start,
		"$options": "i",
	}

	if input.SearchInput != nil {
		if input.FilterType != nil {
			field := *input.FilterType
			if field == "technician" || field == "patient_name" || field == "insurances" {
				filters[field] = bson.M{
					"$regex":   *input.SearchInput,
					"$options": "i",
				}
			} else {
				filters["$or"] = bson.A{
					bson.M{"technician": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
					bson.M{"patient_name": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
					bson.M{"procedure": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				}
			}
		} else {
			filters["$or"] = bson.A{
				bson.M{"technician": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"patient_name": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"procedure": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
			}
		}
	}

	return filters
}

package appointment_code_mongo_repository

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	resetPasswordCodeEntity := AppointmentModel{
		UserUuid:    input.UserUuid,
		Uuid:        input.Uuid,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		PatientUuid: input.PatientUuid,
		Insurance:   input.Insurance,
		Technician:  input.Technician,
		Location:    input.Location,
		Status:      input.StartDate,
		Procedure:   input.Procedure,
	}

	_, err := f.collection.InsertOne(ctx, resetPasswordCodeEntity)

	if err != nil {
		return err
	}

	return nil
}

func (f *AppointmentRepository) List(ctx context.Context, input dto.ListAppointmentInputDto) ([]entity.Appointment, error) {
	filters := bson.M{}

	var start, end time.Time
	if input.Date == nil {
		now := time.Now()
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		end = start.Add(24 * time.Hour)
	} else {
		start = time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, input.Date.Location())
		end = start.Add(24 * time.Hour)
	}

	filters["start_date"] = bson.M{
		"$gte": start,
		"$lt":  end,
	}

	if input.SearchInput != nil {
		if input.FilterType != nil {
			filters[*input.FilterType] = bson.M{
				"$regex":   *input.SearchInput,
				"$options": "i",
			}
		} else {
			filters["$or"] = bson.A{
				bson.M{"technician": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"patient_name": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"insurances": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
			}
		}
	}

	cursor, err := f.collection.Find(ctx, filters)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var appointments []entity.Appointment
	if err := cursor.All(ctx, &appointments); err != nil {
		return nil, err
	}

	return appointments, nil
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

func (f *AppointmentRepository) Delete(ctx context.Context, input dto.DeleteAppointmentInputDto) {
	filter := bson.M{"uuid": input.Uuid}

	f.collection.FindOneAndDelete(ctx, filter)
}

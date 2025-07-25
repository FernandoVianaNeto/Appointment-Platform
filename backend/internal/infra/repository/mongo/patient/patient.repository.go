package patient_mongo_repository

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PatientRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewPatientRepository(db *mongo.Database) domain_repository.PatientRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.PatientCollection)

	return &PatientRepository{
		db:         db,
		collection: collection,
	}
}

func (f *PatientRepository) Create(ctx context.Context, input entity.Patient) error {
	patientEntity := PatientModel{
		Uuid:         input.Uuid,
		UserUuid:     input.UserUuid,
		Insurance:    input.Insurance,
		Name:         input.Name,
		Phone:        input.Phone,
		Address:      input.Address,
		Email:        input.Email,
		ReminderSent: false,
	}

	_, err := f.collection.InsertOne(ctx, patientEntity)

	return err
}

func (f *PatientRepository) List(ctx context.Context, input dto.ListPatientInputDto) ([]entity.Patient, error) {
	filters := buildListFilters(input)

	limit := int64(domain_response.DEFAULT_ITEMS_PER_PAGE)
	skip := int64((input.Page - 1)) * limit

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetSort(bson.M{"created_at": -1})

	cursor, err := f.collection.Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var patients []PatientModel
	if err := cursor.All(ctx, &patients); err != nil {
		return nil, err
	}

	patientsEntity := make([]entity.Patient, 0, len(patients))
	for _, patient := range patients {
		patientsEntity = append(patientsEntity, entity.Patient{
			Uuid:      patient.Uuid,
			UserUuid:  patient.UserUuid,
			Name:      patient.Name,
			Phone:     patient.Phone,
			Insurance: patient.Insurance,
			Address:   patient.Address,
			Email:     patient.Email,
		})
	}

	return patientsEntity, nil
}

func (f *PatientRepository) Edit(ctx context.Context, input dto.EditPatientInputDto) error {
	updateFields := bson.M{}

	if input.Name != nil {
		updateFields["name"] = input.Name
	}
	if input.Address != nil {
		updateFields["address"] = input.Address
	}
	if input.Email != nil {
		updateFields["email"] = input.Email
	}
	if input.Phone != nil {
		updateFields["phone"] = input.Phone
	}

	if len(updateFields) == 0 {
		return nil
	}

	update := bson.M{"$set": updateFields}

	filter := bson.M{"user_uuid": input.Uuid, "uuid": input.PatientUuid}

	err := f.collection.FindOneAndUpdate(ctx, filter, update).Err()
	return err
}

func (f *PatientRepository) Delete(ctx context.Context, uuid string) {
	filter := bson.M{"uuid": uuid}

	f.collection.FindOneAndDelete(ctx, filter)
}

func (f *PatientRepository) DeleteMany(ctx context.Context, ids []string) error {
	filter := bson.M{
		"uuid": bson.M{
			"$in": ids,
		},
	}

	_, err := f.collection.DeleteMany(ctx, filter)
	return err
}

func (f *PatientRepository) GetByUuid(ctx context.Context, uuid string) (entity.Patient, error) {
	var model PatientModel

	filter := bson.M{"uuid": uuid}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		return entity.Patient{}, err
	}

	response := entity.Patient{
		Uuid:      model.Uuid,
		Name:      model.Name,
		Phone:     model.Phone,
		Insurance: model.Insurance,
		Address:   model.Address,
		Email:     model.Email,
	}

	return response, err
}

func (f *PatientRepository) CountDocuments(ctx context.Context, input dto.ListPatientInputDto) (int64, error) {
	filters := buildListFilters(input)

	total, err := f.collection.CountDocuments(ctx, filters)

	return total, err
}

func buildListFilters(input dto.ListPatientInputDto) bson.M {
	var filters = bson.M{
		"user_uuid": input.UserUuid,
	}

	if input.Name != nil {
		filters["name"] = bson.M{"$regex": *input.Name, "$options": "i"}
	}

	if input.Uuid != nil {
		filters["uuid"] = bson.M{"$regex": *input.Uuid, "$options": "i"}
	}

	if input.SearchInput != nil {
		if input.FilterType != nil {
			filters[*input.FilterType] = bson.M{
				"$regex":   *input.SearchInput,
				"$options": "i",
			}
		} else {
			filters["$or"] = bson.A{
				bson.M{"phone": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"name": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
				bson.M{"email": bson.M{"$regex": *input.SearchInput, "$options": "i"}},
			}
		}
	}

	return filters
}

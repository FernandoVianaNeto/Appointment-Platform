package patient_mongo_repository

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		Uuid:      input.Uuid,
		Insurance: input.Insurance,
		Name:      input.Name,
		Phone:     input.Phone,
		Address:   input.Address,
		Email:     input.Email,
	}

	_, err := f.collection.InsertOne(ctx, patientEntity)

	return err
}

func (f *PatientRepository) List(ctx context.Context, input dto.ListPatientInputDto) ([]entity.Patient, error) {
	filters := bson.M{}

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

	cursor, err := f.collection.Find(ctx, filters)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var patients []entity.Patient
	if err := cursor.All(ctx, &patients); err != nil {
		return nil, err
	}

	return patients, nil
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

	filter := bson.M{"uuid": input.Uuid}

	err := f.collection.FindOneAndUpdate(ctx, filter, update).Err()
	return err
}

func (f *PatientRepository) Delete(ctx context.Context, input dto.DeletePatientInputDto) {
	filter := bson.M{"uuid": input.Uuid}

	f.collection.FindOneAndDelete(ctx, filter)
}

package mongo_repository

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain_repository.UserRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.UserCollection)

	return &UserRepository{
		db:         db,
		collection: collection,
	}
}

func (f *UserRepository) Create(ctx context.Context, input entity.User) error {
	var passwordString *string

	if input.Password != nil {
		passwordString = new(string)
		*passwordString = string(*input.Password)
	}

	_, err := f.collection.InsertOne(ctx, UserModel{
		Uuid:         input.Uuid,
		Email:        input.Email,
		Name:         input.Name,
		Password:     passwordString,
		AuthProvider: input.AuthProvider,
		GoogleSub:    input.GoogleSub,
	})

	return err
}

func (f *UserRepository) GetByUuid(ctx context.Context, userUuid string) (*entity.User, error) {
	var model UserModel

	filter := bson.M{
		"uuid": userUuid,
	}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	entity := entity.User{
		Uuid:         model.Uuid,
		Email:        model.Email,
		Name:         model.Name,
		AuthProvider: model.AuthProvider,
	}

	return &entity, err
}

func (f *UserRepository) GetByEmailAndAuthProvider(ctx context.Context, email string, authProvider string) (*entity.User, error) {
	var model UserModel

	filter := bson.M{
		"email":         email,
		"auth_provider": authProvider,
	}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	entity := entity.User{
		Uuid:  model.Uuid,
		Email: model.Email,
		Name:  model.Name,
		Password: func() *[]byte {
			if model.Password == nil {
				return nil
			}
			b := []byte(*model.Password)
			return &b
		}(),
		AuthProvider: model.AuthProvider,
	}

	return &entity, err
}

func (f *UserRepository) UpdateByUuid(ctx context.Context, input dto.UpdateUserInputDto) error {
	filter := bson.M{
		"uuid": input.Uuid,
	}
	setFields := bson.M{}

	if input.Email != nil {
		setFields["email"] = *input.Email
	}
	if input.Name != nil {
		setFields["name"] = *input.Name
	}

	fmt.Println(setFields, len(setFields))

	if len(setFields) > 0 {
		update := bson.M{
			"$set": setFields,
		}
		_, err := f.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserRepository) UpdatePassword(ctx context.Context, input dto.UserResetPasswordInputDto) error {
	updateUser := bson.M{}

	filter := bson.M{
		"uuid": input.Uuid,
	}

	if input.NewPassword != nil {
		updateUser["$set"] = bson.M{
			"password": input.NewPassword,
		}
	}

	if len(updateUser) > 0 {
		_, err := f.collection.UpdateOne(ctx, filter, updateUser)
		if err != nil {
			return err
		}
	}

	return nil
}

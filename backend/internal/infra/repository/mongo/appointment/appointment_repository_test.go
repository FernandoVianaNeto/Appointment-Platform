package appointment_mongo_repository_test

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"appointment-platform-backend-backend/internal/domain/dto"
// 	"appointment-platform-backend-backend/internal/domain/entity"
// 	appointment_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/appointment"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// func TestAppointmentRepository_Create(t *testing.T) {
// 	m := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer m.Close()

// 	m.Run("should create appointment successfully", func(mt *mtest.T) {
// 		repo := appointment_mongo_repository.NewAppointmentRepository(mt.DB)

// 		input := dto.CreateAppointmentInputDto{
// 			Uuid:        "123",
// 			UserUuid:    "user-uuid",
// 			StartDate:   time.Now().String(),
// 			EndDate:     time.Now().Add(time.Hour).String(),
// 			PatientUuid: "patient-uuid",
// 			Insurance:   "insurance",
// 			Technician:  "Dr. John",
// 			Location:    "room-1",
// 			Procedure:   "examination",
// 		}

// 		mt.AddMockResponses(mtest.CreateSuccessResponse())

// 		err := repo.Create(context.TODO(), input)
// 		assert.NoError(t, err)
// 	})
// }
// func TestAppointmentRepository_List(t *testing.T) {
// 	m := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer m.Close()

// 	m.Run("should list appointments", func(mt *mtest.T) {
// 		repo := appointment_mongo_repository.NewAppointmentRepository(mt.DB)

// 		// Mock response
// 		doc := bson.D{
// 			{"uuid", "123"},
// 			{"user_uuid", "user-uuid"},
// 			{"start_date", "2025-07-28T10:00"},
// 			{"end_date", "2025-07-28T11:00"},
// 			{"patient_uuid", "patient-uuid"},
// 			{"insurance", "insurance-x"},
// 			{"technician", "tech1"},
// 			{"location", "room1"},
// 			{"status", "scheduled"},
// 			{"procedure", "exam-x"},
// 		}

// 		cursor := mtest.CreateCursorResponse(1, "test.appointments", mtest.FirstBatchResponse(doc))
// 		mt.AddMockResponses(cursor)

// 		result, err := repo.List(context.TODO(), getFakeListInput())
// 		assert.NoError(t, err)
// 		assert.Len(t, result, 1)
// 		assert.Equal(t, "123", result[0].Uuid)
// 	})
// }

// func getFakeAppointment() entity.Appointment {
// 	return entity.Appointment{
// 		Uuid:        "123",
// 		UserUuid:    "user-uuid",
// 		StartDate:   "2025-07-28T10:00",
// 		EndDate:     "2025-07-28T11:00",
// 		PatientUuid: "patient-uuid",
// 		Insurance:   "insurance-x",
// 		Technician:  "tech1",
// 		Location:    "room1",
// 		Status:      "scheduled",
// 		Procedure:   "exam-x",
// 	}
// }

// func getFakeListInput() dto.ListAppointmentInputDto {
// 	date := "2025-07-28"
// 	return dto.ListAppointmentInputDto{
// 		UserUuid: "user-uuid",
// 		Page:     1,
// 		Date:     &date,
// 	}
// }

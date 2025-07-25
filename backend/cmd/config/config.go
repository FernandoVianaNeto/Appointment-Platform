package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApplicationCfg  *ApplicationConfig
	MongoCfg        *MongoConfig
	GoogleAuthCfg   *GoogleAuthConfig
	SendGridCfg     *SendGridConfig
	ReminderCronCfg *ReminderCronConfig
)

const (
	AppName     = "appointment-platform-backend-backend"
	AppVersion  = "1.0.0"
	Development = "development"
	Staging     = "stage"
	Production  = "production"
)

type ApplicationConfig struct {
	Env         string
	AppVersion  string
	AppPort     int
	JwtSecret   string
	Environment string
}

type GoogleAuthConfig struct {
	ClientId string
}

type SendGridConfig struct {
	ApiKey string
}

type ReminderCronConfig struct {
	Window int
}

type MinIOConfig struct {
	Host                   string
	User                   string
	Password               string
	ProfileBucket          string
	PresignedURLExpiration string
}

type NatsConfig struct {
	Host      string
	User      string
	Password  string
	UserTopic string
}

type MongoConfig struct {
	UserCollection              string
	AppointmentCollection       string
	PatientCollection           string
	ResetPasswordCodeCollection string
	Dsn                         string
	Database                    string
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitializeConfigs() {
	initialize()
	initializeApplicationConfigs()
	initializeMongoConfigs()
	initializeGoogleAuthConfigs()
	initializeSendGridConfigs()
	initializeReminderCronConfig()
}

func getEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &ApplicationConfig{
			Env:        getEnv("APP_ENV", "local"),
			AppVersion: AppVersion,
			AppPort:    getEnvAsInt("APP_PORT", 3001),
			JwtSecret:  getEnv("JWT_SECRET", "random_hash"),
		}
	}
}

func initializeMongoConfigs() {
	if MongoCfg == nil {
		MongoCfg = &MongoConfig{
			UserCollection:              getEnv("MONGO_USER_COLLECTION", "user"),
			ResetPasswordCodeCollection: getEnv("MONGO_RESET_PASSWORD_CODE_COLLECTION", "reset_password_code"),
			AppointmentCollection:       getEnv("APPOINTMENT_COLLECTION", "appointment"),
			PatientCollection:           getEnv("PATIENT_COLLECTION", "patient"),
			Dsn:                         getEnv("MONGO_DSN", "mongodb://localhost:27017"),
			Database:                    getEnv("MONGO_DB", "appointment-plataform"),
		}
	}
}

func initializeGoogleAuthConfigs() {
	if GoogleAuthCfg == nil {
		GoogleAuthCfg = &GoogleAuthConfig{
			ClientId: getEnv("GOOGLE_CLIENT_ID", "your-google-client-id"),
		}
	}
}

func initializeSendGridConfigs() {
	if SendGridCfg == nil {
		SendGridCfg = &SendGridConfig{
			ApiKey: getEnv("SEND_GRID_API_KEY", "your-sendgrid-api-key"),
		}
	}
}

func initializeReminderCronConfig() {
	if ReminderCronCfg == nil {
		ReminderCronCfg = &ReminderCronConfig{
			Window: getEnvAsInt("REMINDER_CRON_WINDOW", 24),
		}
	}
}

package cli

import (
	configs "appointment-platform-backend-backend/cmd/config"
	app "appointment-platform-backend-backend/internal/application"
	appointment_usecase "appointment-platform-backend-backend/internal/application/usecase/appointment"
	"appointment-platform-backend-backend/internal/infra/adapter/sendgrid"
	"appointment-platform-backend-backend/internal/infra/cron"
	appointment_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/appointment"
	mongoPkg "appointment-platform-backend-backend/pkg/mongo"
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(appointmentReminderCronCmd)
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root - main command application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Initialize Http Server",
	Run: func(cmd *cobra.Command, args []string) {
		port := configs.ApplicationCfg.AppPort
		if port == 0 {
			os.Exit(1)
		}

		srv := app.NewApplication()
		if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
			os.Exit(1)
		}
	},
}

var appointmentReminderCronCmd = &cobra.Command{
	Use:   "appointment-reminder-cron",
	Short: "Runs a cronjob to reminder the patient about the appointment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		mongoConnectionInput := mongoPkg.MongoInput{
			DSN:      configs.MongoCfg.Dsn,
			Database: configs.MongoCfg.Database,
		}

		db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

		appointmentRepository := appointment_mongo_repository.NewAppointmentRepository(db)

		emailSenderAdapter := sendgrid.NewEmailSenderAdapter(ctx)

		usecase := appointment_usecase.NewGetNextAppointmentsAndSendReminder(appointmentRepository, emailSenderAdapter)

		cron.StartReminderScheduler(ctx, usecase, appointmentRepository, emailSenderAdapter)
	},
}

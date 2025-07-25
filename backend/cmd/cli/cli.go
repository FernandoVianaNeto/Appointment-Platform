package cli

import (
	configs "appointment-platform-backend-backend/cmd/config"
	app "appointment-platform-backend-backend/internal/application"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(appointmentReminderCmd)
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

var appointmentReminderCmd = &cobra.Command{
	Use:   "appointment-reminder-cron",
	Short: "Runs a cronjob to reminder the patient about the appointment",
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

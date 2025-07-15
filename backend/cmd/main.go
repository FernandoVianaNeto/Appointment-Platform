package main

import (
	"appointment-platform-backend-backend/cmd/cli"
	configs "appointment-platform-backend-backend/cmd/config"
)

func main() {
	configs.InitializeConfigs()

	cli.Execute()
}

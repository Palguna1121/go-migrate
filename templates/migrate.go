package templates

var MigrateTemplate = `package main

import (
	"github.com/Palguna1121/go-migrate/cmd"
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
	_ "%[2]s/migrations"
	localconfig "your_project/config" // Replace with your actual config package
	"strconv"
)

func init() {
	port, err := strconv.Atoi(localconfig.ENV.DB_PORT)
	if err != nil {
		panic("Invalid DB_PORT in environment variables")
	}
	config.Config = config.DatabaseConfig{
		Host:     localconfig.ENV.DB_HOST,
		Port:     port,
		Username: localconfig.ENV.DB_USER,
		Password: localconfig.ENV.DB_PASSWORD,
		Dbname:   localconfig.ENV.DB_NAME,
	}

	config.Migrator = %[1]s.InitMigrator()
	config.Driver = "%[1]s"
}

func main() {
	cmd.Execute()
}
`

package templates

var MigrateTemplate = `package main

import (
	"go-migrate/cmd"
	"go-migrate/config"
	"go-migrate/pkg/lib/%[1]s"
	_ "%[2]s/migrations"
)

func init() {
	config.Config = config.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "",
		Dbname:   "test",
	}

	config.Migrator = %[1]s.InitMigrator()
	config.Driver = "%[1]s"
}

func main() {
	cmd.Execute()
}
`

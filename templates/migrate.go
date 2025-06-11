package templates

var MigrateTemplate = `package main

import (
	"github.com/Palguna1121/go-migrate/cmd"
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
	_ "%[2]s/migrations"
	"strconv"
)

func init() {
	config.Config = config.DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "",
		Dbname:   "test_your_db",
	}

	config.Migrator = %[1]s.InitMigrator()
	config.Driver = "%[1]s"
}

func main() {
	cmd.Execute()
}
`

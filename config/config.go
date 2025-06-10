package config

import "go-migrate/pkg/interfaces"

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

var (
	Config     DatabaseConfig
	Migrations []interfaces.Migration
	Migrator   interfaces.Migrator
	Driver     string
)

package config

import "github.com/Palguna1121/go-migrate/pkg/interfaces"

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

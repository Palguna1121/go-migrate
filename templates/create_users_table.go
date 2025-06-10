package templates

var UserMigrationTemplate = `package migrations

import (
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
)

func init() {
	config.Migrations = append(config.Migrations, CreateUsersTable())
}

type UsersTable struct{}

func CreateUsersTable() interfaces.Migration {
	return &UsersTable{}
}

func (t *UsersTable) Up() error {
	// Buat tabel users
	if err := mysql.Schema.Create("users", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("name", 50)
		table.String("email", 50).Unique()
		table.DateTime("email_verified_at").Nullable()
		table.String("password", 255)
		table.String("remember_token", 100).Nullable()
		table.Timestamps()
		table.DateTime("deleted_at").Nullable()
	}); err != nil {
		return err
	}

	// Buat tabel password_reset_tokens
	if err := mysql.Schema.Create("password_reset_tokens", func(table interfaces.Blueprint) {
		table.String("email", 255).Primary()
		table.String("token", 255)
		table.Timestamps()
	}); err != nil {
		return err
	}

	// Buat tabel sessions
	if err := mysql.Schema.Create("sessions", func(table interfaces.Blueprint) {
		table.String("id", 255).Primary()
		table.Integer("user_id", 11).Nullable().Index()
		table.String("ip_address", 45).Nullable()
		table.Text("user_agent").Nullable()
		table.Text("payload")
		table.Integer("last_activity", 11).Index()
	}); err != nil {
		return err
	}

	return nil
}

func (t *UsersTable) Down() error {
	// Drop tabel sessions
	if err := mysql.Schema.DropIfExists("sessions"); err != nil {
		return err
	}

	// Drop tabel password_reset_tokens
	if err := mysql.Schema.DropIfExists("password_reset_tokens"); err != nil {
		return err
	}

	// Drop tabel users
	if err := mysql.Schema.DropIfExists("users"); err != nil {
		return err
	}

	return nil
}
`

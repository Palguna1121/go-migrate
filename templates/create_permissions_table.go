package templates

var PermissionsMigrationTemplate = `package migrations

import (
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
)

func init() {
	config.Migrations = append(config.Migrations, CreatePermissionTable())
}

type PermissionTable struct{}

func CreatePermissionTable() interfaces.Migration {
	return &PermissionTable{}
}

func (t *PermissionTable) Up() error {
	if err := mysql.Schema.Create("permissions", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("name", 255)
		table.String("guard_name", 255)
		table.Timestamps()

		// Unique constraint
		table.Unique("name,guard_name")
	}); err != nil {
		return err
	}

	if err := mysql.Schema.Create("roles", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("name", 255)
		table.String("guard_name", 255)
		table.Timestamps()

		// Unique constraint
		table.Unique("name", "guard_name")
	}); err != nil {
		return err
	}

	// Create 'model_has_permissions' table
	if err := mysql.Schema.Create("model_has_permissions", func(table interfaces.Blueprint) {
		table.Foreign("permission_id").Reference("id").On("permissions").OnDelete("cascade")
		table.String("model_type", 255)
		table.Integer("model_id", 11)

		// Corrected Index syntax for multiple columns
		table.Index("model_id", "model_type") // Assuming it takes multiple string arguments

		table.Primary("permission_id", "model_id", "model_type")
	}); err != nil {
		return err
	}

	// Create 'model_has_roles' table
	if err := mysql.Schema.Create("model_has_roles", func(table interfaces.Blueprint) {
		table.Foreign("role_id").Reference("id").On("roles").OnDelete("cascade")

		table.String("model_type", 255)
		table.Integer("model_id", 11)

		table.Index("model_id", "model_type") // Assuming it takes multiple string arguments
		table.Primary("role_id", "model_id", "model_type")
	}); err != nil {
		return err
	}

	// Create 'role_has_permissions' table
	if err := mysql.Schema.Create("role_has_permissions", func(table interfaces.Blueprint) {
		table.Foreign("permission_id").Reference("id").On("permissions").OnDelete("cascade")
		table.Foreign("role_id").Reference("id").On("roles").OnDelete("cascade")
		table.Primary("permission_id", "role_id")
	}); err != nil {
		return err
	}

	return nil
}

func (t *PermissionTable) Down() error {
	if err := mysql.Schema.DropIfExists("permissions"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("roles"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("model_has_permissions"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("model_has_roles"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("role_has_permissions"); err != nil {
		return err
	}
	return nil
}
`

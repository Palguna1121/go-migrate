package templates

var PersonalAccessTokensMigrationTemplate = `package migrations

import (
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
)

func init() {
	config.Migrations = append(config.Migrations, CreatePersonalAccessTokensTable())
}

type PersonalAccessTokensTable struct{}

func CreatePersonalAccessTokensTable() interfaces.Migration {
	return &PersonalAccessTokensTable{}
}

func (t *PersonalAccessTokensTable) Up() error {
	return mysql.Schema.Create("personal_access_tokens", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("tokenable_type", 255)
		table.Integer("tokenable_id", 11).Index()
		table.String("name", 255)
		table.String("token", 64).Unique()
		table.String("abilities", 255).Nullable()
		table.DateTime("last_used_at").Nullable()
		table.DateTime("expires_at").Nullable()
		table.Timestamps()
	})
}

func (t *PersonalAccessTokensTable) Down() error {
	return mysql.Schema.DropIfExists("personal_access_tokens")
}
`

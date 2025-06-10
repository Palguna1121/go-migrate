package templates

var PersonalAccessTokensMigrationTemplate = `package migrations

import (
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
)

func init() {
	config.Migrations = append(config.Migrations, CreateCacheTable())
}

type CacheTable struct{}

func CreateCacheTable() interfaces.Migration {
	return &CacheTable{}
}

func (t *CacheTable) Up() error {
	if err := mysql.Schema.Create("cache", func(table interfaces.Blueprint) {
		table.String("key", 255).Primary()
		table.Text("value")
		table.Integer("expiration", 11).Nullable()
	}); err != nil {
		return err
	}

	if err := mysql.Schema.Create("cache_locks", func(table interfaces.Blueprint) {
		table.String("key", 255).Primary()
		table.String("owner", 255).Nullable()
		table.Integer("expiration", 11).Nullable()
	}); err != nil {
		return err
	}

	return nil
}

func (t *CacheTable) Down() error {
	if err := mysql.Schema.DropIfExists("cache"); err != nil {
		return err
	}

	if err := mysql.Schema.DropIfExists("cache_locks"); err != nil {
		return err
	}

	return nil
}
`

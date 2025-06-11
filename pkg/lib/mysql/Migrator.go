package mysql

import (
	"fmt"
	"strings"

	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/model"
)

type migrator struct{}

func InitMigrator() interfaces.Migrator {
	return &migrator{}
}

func (m *migrator) CheckTable() (bool, error) {
	driver, err := NewDriver()
	if err != nil {
		return false, err
	}
	defer driver.Close()

	var count int64
	sql := "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'migrations'"
	err = driver.QueryRow(&count, sql)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *migrator) CreateTable() error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}
	defer driver.Close()

	sqls := []string{
		"CREATE TABLE `migrations` (`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, `migration` varchar(255) NOT NULL, `batch` int(11) NOT NULL, PRIMARY KEY (`id`));",
	}

	for _, sql := range sqls {
		if err := driver.Execute(sql); err != nil {
			return err
		}
	}

	return nil
}

func (m *migrator) DropTableIfExists() error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}
	defer driver.Close()

	sql := "DROP TABLE IF EXISTS migrations"
	return driver.Execute(sql)
}

func (m *migrator) DropAllTable() error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}
	defer driver.Close()

	var tables []string
	sql := "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_type = 'BASE TABLE'"
	err = driver.Query(&tables, sql)
	if err != nil {
		return err
	}

	if len(tables) == 0 {
		return nil
	}

	err = driver.Execute("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return err
	}

	sql = fmt.Sprintf("DROP TABLE IF EXISTS %s", strings.Join(tables, ","))
	err = driver.Execute(sql)

	driver.Execute("SET FOREIGN_KEY_CHECKS = 1")

	return err
}

func (m *migrator) GetMigrations() ([]model.Migration, error) {
	driver, err := NewDriver()
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	var migrations []model.Migration
	sql := "SELECT id, migration, batch FROM `migrations` ORDER BY id"
	err = driver.Query(&migrations, sql)
	return migrations, err
}

func (m *migrator) WriteRecord(migration string, batch int) error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}
	defer driver.Close()

	sql := "INSERT INTO `migrations`(`migration`, `batch`) VALUES (?, ?)"
	return driver.Execute(sql, migration, batch)
}

func (m *migrator) DeleteRecord(id int) error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}
	defer driver.Close()

	sql := "DELETE FROM `migrations` WHERE id = ?"
	return driver.Execute(sql, id)
}

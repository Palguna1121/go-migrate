package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Palguna1121/response-std/config"
)

type driver struct {
	db *gorm.DB
}

func NewDriver() (*driver, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.ENV.DB_USER,
		config.ENV.DB_PASSWORD,
		config.ENV.DB_HOST,
		config.ENV.DB_PORT,
		config.ENV.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect mysql server failed, err:%v", err)
	}

	return &driver{db: db}, nil
}

// Execute - menjalankan raw SQL tanpa mengembalikan rows
func (d *driver) Execute(sql string, args ...interface{}) error {
	return d.db.Exec(sql, args...).Error
}

// Query - SELECT query yang mengembalikan multiple rows
func (d *driver) Query(dest interface{}, sql string, args ...interface{}) error {
	return d.db.Raw(sql, args...).Scan(dest).Error
}

// QueryRow - SELECT query single row
func (d *driver) QueryRow(dest interface{}, sql string, args ...interface{}) error {
	return d.db.Raw(sql, args...).Scan(dest).Error
}

// Select - alias Query
func (d *driver) Select(dest interface{}, sql string, args ...interface{}) error {
	return d.Query(dest, sql, args...)
}

// Close - menutup koneksi database
func (d *driver) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

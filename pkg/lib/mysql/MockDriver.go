package mysql

import (
	"gorm.io/gorm"
)

// MockDriver simulates a GORM-like DB interaction
type MockDriver struct {
	sqls    []string
	isClose bool
}

// NewMockDriver creates a new mock driver
func NewMockDriver() (*MockDriver, error) {
	return &MockDriver{isClose: false}, nil
}

// Create simulates an insert operation
func (d *MockDriver) Create(value interface{}) *gorm.DB {
	d.sqls = append(d.sqls, "CREATE SQL") // You can make it dynamic if needed
	return &gorm.DB{}
}

// Find simulates a select query
func (d *MockDriver) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	d.sqls = append(d.sqls, "FIND SQL")
	return &gorm.DB{}
}

// Raw simulates raw SQL query execution
func (d *MockDriver) Raw(sql string, values ...interface{}) *gorm.DB {
	d.sqls = append(d.sqls, sql)
	return &gorm.DB{}
}

// Close marks the mock as closed
func (d *MockDriver) Close() error {
	d.isClose = true
	return nil
}

// GetSqls returns all the executed SQL statements
func (d *MockDriver) GetSqls() []string {
	return d.sqls
}

// IsClose returns true if the mock was closed
func (d *MockDriver) IsClose() bool {
	return d.isClose
}

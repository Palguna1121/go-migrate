package interfaces

type Driver interface {
	Execute(sql string, args ...interface{}) error
	Query(dest interface{}, sql string, args ...interface{}) error
	Select(dest interface{}, sql string, args ...interface{}) error
	Close() error
}

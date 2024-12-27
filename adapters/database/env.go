package database

type Env interface {
	GetDBType() string
	GetDBPath() string
}

package database

import (
	"web1/domain/use-cases"
)

type Database interface {
	usecases.Database
	Close() error
}

func InitializeDatabase(env Env) Database {
	switch env.GetDBType() {
	case "inmemory":
		db, _ := NewInMemoryDatabase()
		return db
	case "sqlite", "sqlite3":
		db, err := NewSQLiteDatabase(env.GetDBPath())
		if err != nil {
			panic(err)
		}
		return db
	default:
		panic("Unknown database type")
	}
}

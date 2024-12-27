package database

import (
	"web1/domain/interfaces"
)

func InitializeDatabase(env interfaces.IEnv) interfaces.IDatabase {
	switch env.GetDBType() {
	case "inmemory":
		db, _ := NewInMemoryDatabase()
		return db
	case "sqlite", "sqlite3":
		// Currently the database file is hardcoded
		db, err := NewSQLiteDatabase("file:db.sqlite3")
		if err != nil {
			panic(err)
		}
		return db
	default:
		panic("Unknown database type")
	}
}

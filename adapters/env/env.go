package environment

import (
	"os"
	"strconv"
)

const (
	DEFAULT_PORT    = 8080
	DEFAULT_DB_TYPE = "inmemory"
	DEFAULT_DB_PATH = "file:db.sqlite3"
)

type Env struct {
	serverPort string
	dbType     string
	dbPath     string
}

func NewEnv() *Env {
	return &Env{
		serverPort: os.Getenv("SERVER_PORT"),
		dbType:     os.Getenv("DB_TYPE"),
	}
}

func (e *Env) GetServerPort() int {
	port, err := strconv.Atoi(e.serverPort)

	if err != nil {
		return DEFAULT_PORT
	}

	return port
}

func (e *Env) GetDBType() string {
	if e.dbType == "" {
		return DEFAULT_DB_TYPE
	}

	return e.dbType
}

// Currently always return hardcoded value
func (e *Env) GetDBPath() string {
	return DEFAULT_DB_PATH
}

package environment

import (
	"os"
	"strconv"
)

const (
	DEFAULT_PORT    = 8080
	DEFAULT_DB_TYPE = "inmemory"
)

type Env struct {
	serverPort string
	dbType     string
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

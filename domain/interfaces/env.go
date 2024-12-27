package interfaces

type IEnv interface {
	GetServerPort() int
	GetDBType() string
}

package config

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvTest        = "test"

	ServerListen       = "0.0.0.0"
	ServerPort         = "8000"
	ServerTimeoutRead  = "15s"
	ServerTimeoutWrite = "15s"
	ServerTimeoutIdle  = "60s"

	MariadbHost = "127.0.0.1"
	MariadbPort = "3306"

	ReadBufferSize = 8192
)

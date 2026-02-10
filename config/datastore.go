package config

type mariadbCfg struct {
	Host      string
	Port      string
	User      string
	Password  string
	Database  string
	Migration bool
}

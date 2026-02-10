package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tomatosAt/IT01-api/pkg/util"
)

type Config struct {
	App    appConfig
	Secret secretCfg
	Server serverCfg
	Syslog syslogCfg
	DBMain mariadbCfg // Galera Cluster
	DBLog  mariadbCfg // Standalone
}

type appConfig struct {
	Name       string
	Version    string
	Mode       string
	PrefixPath string
	StorageDir string
	LogLevel   logrus.Level
	Url        string
}

type secretCfg struct {
	EncryptKey string
}

func (c *appConfig) IsDebug() bool {
	return c.LogLevel == logrus.DebugLevel
}

func LoadConfig(file string, version string) *Config {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalln("load config file error:", err.Error())
	}
	// Set default configuration
	viper.SetDefault("app.name", "It")
	viper.SetDefault("app.version", version)
	viper.SetDefault("app.mode", EnvDevelopment)
	viper.SetDefault("app.prefix.path", "/api")
	viper.SetDefault("app.log.level", "info")
	viper.SetDefault("app.url", "")

	// Secret
	viper.SetDefault("secret.encrypt.key", "")

	// Server
	viper.SetDefault("server.listen", ServerListen)
	viper.SetDefault("server.port", ServerPort)
	viper.SetDefault("server.timeout.read", ServerTimeoutRead)
	viper.SetDefault("server.timeout.write", ServerTimeoutWrite)
	viper.SetDefault("server.timeout.idle", ServerTimeoutIdle)
	viper.SetDefault("server.header", viper.GetString("app.name"))
	viper.SetDefault("server.proxy.header", fiber.HeaderXForwardedFor)
	viper.SetDefault("server.enable.cors", "false")
	viper.SetDefault("server.buffer.read", ReadBufferSize)
	viper.SetDefault("server.body.limit", 10*1024*1024*1024) // 10GB

	// DBMain
	viper.SetDefault("db.main.host", MariadbHost)
	viper.SetDefault("db.main.port", MariadbPort)
	viper.SetDefault("db.main.username", "")
	viper.SetDefault("db.main.password", "")
	viper.SetDefault("db.main.database", "")
	viper.SetDefault("db.main.migration", false)

	config := &Config{
		App: appConfig{
			Name:       viper.GetString("app.name"),
			Version:    viper.GetString("app.version"),
			Mode:       viper.GetString("app.mode"),
			PrefixPath: viper.GetString("app.prefix.path"),
			StorageDir: viper.GetString("app.storage"),
			Url:        viper.GetString("app.url"),
		},
		Secret: secretCfg{
			EncryptKey: viper.GetString("secret.encrypt.key"),
		},
		Server: serverCfg{
			ListenIp:       viper.GetString("server.listen"),
			Port:           viper.GetString("server.port"),
			TimeoutRead:    util.ParseDuration(viper.GetString("server.timeout.read")),
			TimeoutWrite:   util.ParseDuration(viper.GetString("server.timeout.write")),
			TimeoutIdle:    util.ParseDuration(viper.GetString("server.timeout.idle")),
			ServerHeader:   viper.GetString("server.header"),
			ProxyHeader:    viper.GetString("server.proxy.header"),
			EnableCORS:     viper.GetBool("server.enable.cors"),
			ReadBufferSize: viper.GetInt("server.buffer.read"),
			BodyLimit:      viper.GetInt("server.body_limit"),
		},
		DBMain: mariadbCfg{
			Host:      viper.GetString("db.main.host"),
			Port:      viper.GetString("db.main.port"),
			User:      viper.GetString("db.main.username"),
			Password:  viper.GetString("db.main.password"),
			Database:  viper.GetString("db.main.database"),
			Migration: viper.GetBool("db.main.migration"),
		},
	}
	config.App.LogLevel, err = logrus.ParseLevel(viper.GetString("app.log.level"))
	if err != nil {
		config.App.LogLevel = logrus.InfoLevel
	}
	return config
}

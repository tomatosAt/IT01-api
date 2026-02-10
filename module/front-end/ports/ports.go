package ports

import (
	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/IT01-api/config"
	"github.com/tomatosAt/IT01-api/pkg/database"
)

type Repository interface {
	Module() string
	AppCfg() *config.Config
	Log() *logrus.Entry
	DB() *database.Client
}

type Service interface {
}

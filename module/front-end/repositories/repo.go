package repositories

import (
	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/IT01-api/app"
	"github.com/tomatosAt/IT01-api/config"
	"github.com/tomatosAt/IT01-api/pkg/database"
	"github.com/tomatosAt/IT01-api/pkg/requests"
)

const moduleName = "front-end"

type Repository struct {
	app    *app.Context
	http   *requests.HttpClient
	log    *logrus.Entry
	dbMain *database.Client
}

func (r Repository) AppCfg() *config.Config {
	return r.app.Config
}

func (r Repository) Module() string {
	return moduleName
}

func (r Repository) Log() *logrus.Entry {
	return r.log.Dup()
}

func (r Repository) DB() *database.Client {
	return r.dbMain
}

func New(app *app.Context) (*Repository, error) {
	l := app.NewLogger().WithField("module", moduleName)
	dbMain, err := app.NewDBMainClient(l)
	if err != nil {
		return nil, err
	}
	httpClient := requests.NewHttpClient(app.AddSyslogHook(l, moduleName))
	return &Repository{
		app:    app,
		http:   httpClient,
		log:    l,
		dbMain: dbMain,
	}, nil
}

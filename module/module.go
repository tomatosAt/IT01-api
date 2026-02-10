package module

import (
	"github.com/tomatosAt/IT01-api/app"
	frontend "github.com/tomatosAt/IT01-api/module/front-end"
)

func Create(app *app.Context) error {
	l := app.NewLogger().WithField("module", "generic")

	if err := frontend.Create(app); err != nil {
		l.Errorln("[x] Create FrontEndAPI module error -:", err)
		return err
	}
	return nil
}

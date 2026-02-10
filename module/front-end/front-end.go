package frontend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/IT01-api/app"
	"github.com/tomatosAt/IT01-api/module/front-end/handler"
	"github.com/tomatosAt/IT01-api/module/front-end/repositories"
	service "github.com/tomatosAt/IT01-api/module/front-end/service"
)

func Create(app *app.Context) error {
	repo, err := repositories.New(app)
	if err != nil {
		return err
	}
	// services
	svc := service.New(repo)
	h := handler.NewHandler(svc)
	prefixPath := app.Config.App.PrefixPath + "-fe"
	g := app.Router.Group(prefixPath)
	addRouter(g, h)
	return nil
}

func addRouter(r fiber.Router, h *handler.Handler) {
	v1 := r.Group("/v1")
	preRegist := v1.Group("user")
	preRegist.Post("", h.RegisterUserHandler) //add

}

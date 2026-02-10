package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func (ctx *Context) InitFiberServer() {
	ctx.log.Infoln("[*] Initialize fiber router")
	cfg := fiber.Config{
		ReadTimeout:           ctx.Config.Server.TimeoutRead,
		WriteTimeout:          ctx.Config.Server.TimeoutWrite,
		IdleTimeout:           ctx.Config.Server.TimeoutIdle,
		DisableStartupMessage: true,
		Prefork:               false,
		ServerHeader:          ctx.Config.Server.ServerHeader,
		ProxyHeader:           ctx.Config.Server.ProxyHeader,
		ErrorHandler:          serverErrorHandler,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
		ReadBufferSize:        ctx.Config.Server.ReadBufferSize,
		BodyLimit:             ctx.Config.Server.BodyLimit,
	}
	r := fiber.New(cfg)
	if ctx.Config.Server.EnableCORS {
		ctx.log.Infoln("[*] Used fiber cors middleware")
		r.Use(cors.New())
	}
}

var serverErrorHandler = func(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if code >= fiber.StatusInternalServerError {
		logrus.Errorln("[PANIC] ", fmt.Sprintf("[%s]", ctx.IP()), ctx.Route().Method, ctx.Route().Path, ":", err)
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return ctx.Status(code).JSON(fiber.Map{"error": err.Error()})
}

func (ctx *Context) StartHTTP() error {
	// print debug route
	if ctx.Config.App.IsDebug() {
		fr := ctx.Router.GetRoutes()
		for _, r := range fr {
			ctx.log.Debugln(r.Name, r.Method, r.Path)
		}
	}

	serverShutdown := make(chan os.Signal, 1)
	signal.Notify(serverShutdown, os.Interrupt)
	go func() {
		// Listen for syscall signals for process to interrupt/quit
		_ = <-serverShutdown
		ctx.log.Infoln("[*] Server terminating...")
		if err := ctx.Router.Shutdown(); err != nil {
			ctx.log.Errorln(fmt.Sprintf("[x] Server shutdown failed: %+v", err))
		}
	}()

	// Run the server
	srvBound := fmt.Sprintf(
		"%s:%s",
		ctx.Config.Server.ListenIp,
		ctx.Config.Server.Port,
	)

	ctx.log.Infoln(fmt.Sprintf("[*] Starting server at %s", srvBound))
	err := ctx.Router.Listen(srvBound)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		ctx.log.Errorln("[x] Start server error:", err.Error())
		return err
	}
	return nil
}

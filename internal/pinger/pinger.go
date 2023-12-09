package pinger

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type App struct {
	pingChan chan<- struct{}
}

func Start(addr string, pingChan chan struct{}) error {
	pinger := App{pingChan: pingChan}
	app := fiber.New()
	app.Get("/ping", pinger.pingHandler)
	return app.Listen(addr)
}

func (p *App) pingHandler(ctx *fiber.Ctx) error {
	p.pingChan <- struct{}{}
	return ctx.SendStatus(http.StatusAccepted)
}

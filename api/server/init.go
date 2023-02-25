package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mkdr4/hex/pkg/log"
	"github.com/spf13/viper"
)

func Init() *fiber.App {
	cfg := config()
	app := fiber.New(*cfg)

	return app
}

func config() *fiber.Config {
	return &fiber.Config{}
}

func Run(app *fiber.App) {
	if viper.GetBool("http.opt.gs") {
		go listen(app)
		waitShutdown(app)
	} else {
		listen(app)
	}
}

func listen(app *fiber.App) {
	if viper.GetBool("http.opt.ssl") {
		if err := app.ListenTLS(
			":"+viper.GetString("http.port.ssl"),
			viper.GetString("http.path.cert"),
			viper.GetString("http.path.key")); err != nil {
			log.Err(log.S, err.Error())
		}
	} else {
		if err := app.Listen(":" + viper.GetString("http.port.def")); err != nil {
			log.Err(log.S, err.Error())
		}
	}
}

func waitShutdown(app *fiber.App) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
	log.Info(log.S, "start server shutdown")
	app.Shutdown()
	log.Info(log.S, "server shutdown!")
}

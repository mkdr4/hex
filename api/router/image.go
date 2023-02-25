package router

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/mkdr4/hex/api/controller/image"
)

func setupImageRouter(rs *routerSetup) {
	data := rs.s.Group("/image")

	image(rs, data)
}

func image(rs *routerSetup, r fiber.Router) {
	cc := controller.ImageController(rs.db)

	r.Get("/*", cc.Image)
}

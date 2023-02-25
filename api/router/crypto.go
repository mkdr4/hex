package router

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/mkdr4/hex/api/controller/crypto"
)

func setupDataRouter(rs *routerSetup) {
	data := rs.s.Group("/data")

	cryptocurrencies(rs, data)
	wallets(rs, data)
	networks(rs, data)
}

func cryptocurrencies(rs *routerSetup, r fiber.Router) {
	cc := controller.DataController(rs.db)

	r.Get("/cryptocurrencies", cc.Cryptocurrencies)
}

func wallets(rs *routerSetup, r fiber.Router) {
	wc := controller.DataController(rs.db)

	r.Get("/wallets/:cryptocurrency", wc.Wallets)
}

func networks(rs *routerSetup, r fiber.Router) {
	nc := controller.DataController(rs.db)

	r.Get("/networks/:cryptocurrency", nc.Networks)
}

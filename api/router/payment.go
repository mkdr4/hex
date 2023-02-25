package router

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/mkdr4/hex/api/controller/payment"
	"github.com/mkdr4/hex/api/middleware"
	"github.com/mkdr4/hex/internal/domain"
)

func setupPaymentRouter(rs *routerSetup) {
	pay := rs.s.Group("/payment")

	create(rs, pay)
	info(rs, pay)
	update(rs, pay)
	delete(rs, pay)
}

func create(rs *routerSetup, r fiber.Router) {
	cc := controller.CreateController(rs.db)

	r.Post("/create", middleware.Setup(rs.db, domain.Medium), cc.PaymentCreate)
}

func info(rs *routerSetup, r fiber.Router) {
	ic := controller.InfoController(rs.db)

	r.Get("/:id", middleware.Setup(rs.db, domain.Low), ic.PaymentInfo)
}

func update(rs *routerSetup, r fiber.Router) {
	uc := controller.UpdateController(rs.db)

	r.Post("/:id", middleware.Setup(rs.db, domain.Medium), uc.PaymentUpdate)
}

func delete(rs *routerSetup, r fiber.Router) {
	dc := controller.DeleteController(rs.db)

	r.Delete("/:id", middleware.Setup(rs.db, domain.Medium), dc.PaymentDelete)
}

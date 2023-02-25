package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type routerSetup struct {
	s  *fiber.App
	db *sqlx.DB
}

func Setup(s *fiber.App, db *sqlx.DB) {
	rs := &routerSetup{
		s:  s,
		db: db,
	}

	setupPaymentRouter(rs)
	setupDataRouter(rs)
	setupImageRouter(rs)
}

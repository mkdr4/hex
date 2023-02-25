package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/pkg/database"
)

type paymentDeleteController struct {
	db *sqlx.DB
}

func DeleteController(db *sqlx.DB) *paymentDeleteController {
	return &paymentDeleteController{
		db: db,
	}
}

func (c *paymentDeleteController) PaymentDelete(ctx *fiber.Ctx) error {
	if err := database.DeletePayment(c.db, ctx.Params("id")); err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.SendString("Delete")
}

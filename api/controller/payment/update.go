package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/pkg/database"
)

type paymentUpdateController struct {
	db *sqlx.DB
}

func UpdateController(db *sqlx.DB) *paymentUpdateController {
	return &paymentUpdateController{
		db: db,
	}
}

func (c *paymentUpdateController) PaymentUpdate(ctx *fiber.Ctx) error {
	if err := database.UpdatePayment(c.db, ctx.Body()); err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.SendString("Successfuly update payment")
}

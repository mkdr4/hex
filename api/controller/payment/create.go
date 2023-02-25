package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/pkg/database"
)

type paymentCreateController struct {
	db *sqlx.DB
}

func CreateController(db *sqlx.DB) *paymentCreateController {
	return &paymentCreateController{
		db: db,
	}
}

func (c *paymentCreateController) PaymentCreate(ctx *fiber.Ctx) error {
	if err := database.CreatePayment(c.db, ctx.Body()); err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.SendString("Succesfuly create payment")
}

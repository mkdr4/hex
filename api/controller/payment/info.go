package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/pkg/database"
)

type paymentInfoController struct {
	db *sqlx.DB
}

func InfoController(db *sqlx.DB) *paymentInfoController {
	return &paymentInfoController{
		db: db,
	}
}

func (ic *paymentInfoController) PaymentInfo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	payment, err := database.GetPayment(ic.db, id)
	if err != nil {
		return ctx.SendString(err.Error())
	} else if len(payment) == 0 {
		return ctx.SendString("Not found payment")
	}

	b, _ := json.MarshalIndent(payment[0], "", "  ")

	return ctx.SendString(string(b))
}

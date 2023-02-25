package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/pkg/database"
)

type dataController struct {
	db *sqlx.DB
}

func DataController(db *sqlx.DB) *dataController {
	return &dataController{
		db: db,
	}
}

func (c *dataController) Cryptocurrencies(ctx *fiber.Ctx) error {
	crypto, err := database.GetCryptocurrencies(c.db)
	if err != nil {
		return ctx.SendString(err.Error())
	}

	b, _ := json.MarshalIndent(crypto, "", "  ")
	return ctx.SendString(string(b))
}

func (c *dataController) Wallets(ctx *fiber.Ctx) error {
	cryptoName := ctx.Params("cryptocurrency")
	wallets, err := database.GetWallets(c.db, cryptoName)
	if err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.SendString(string(*wallets))
}

func (c *dataController) Networks(ctx *fiber.Ctx) error {
	cryptoName := ctx.Params("cryptocurrency")
	networks, err := database.GetNetworks(c.db, cryptoName)
	if err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.SendString(string(*networks))
}

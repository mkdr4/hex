package controller

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type imageController struct {
	db *sqlx.DB
}

func ImageController(db *sqlx.DB) *imageController {
	return &imageController{
		db: db,
	}
}

func (c *imageController) Image(ctx *fiber.Ctx) error {
	imageName := ctx.Params("*")
	if _, err := os.Stat(fmt.Sprintf("internal/image/%s.svg", imageName)); err != nil {
		ctx.Status(404)
		return ctx.SendString("No such image")
	}

	return ctx.SendFile(fmt.Sprintf("internal/image/%s.svg", imageName))
}

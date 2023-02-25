package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/jmoiron/sqlx"
)

type middlewareController struct {
	db *sqlx.DB
}

func Setup(db *sqlx.DB, typeAuth int8) func(*fiber.Ctx) error {
	cfg := config(db, typeAuth)
	auth := keyauth.New(*cfg)

	return auth
}

func config(db *sqlx.DB, ta int8) *keyauth.Config {
	ac := &middlewareController{db: db}
	auth := typeAuth(ac, ta)
	return &keyauth.Config{
		Validator: auth,
		KeyLookup: "header:Authorization",
	}
}

func typeAuth(cc *middlewareController, mode int8) func(c *fiber.Ctx, s string) (bool, error) {
	switch mode {
	case 3:
		return cc.adminAuth

	case 2:
		return cc.storeAuth

	default:
		return cc.basicAuth
	}
}

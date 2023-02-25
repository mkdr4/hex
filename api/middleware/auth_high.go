package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/mkdr4/hex/internal/domain"
	"github.com/mkdr4/hex/pkg/database"
	"github.com/spf13/viper"
)

func (c *middlewareController) adminAuth(ctx *fiber.Ctx, key string) (bool, error) {
	if !viper.GetBool("httpserver.auth") {
		return true, nil
	}

	auth, err := database.GetAuthkey(c.db, domain.High)
	if err != nil {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}

	for _, a := range auth {
		if a.Key == key {
			return true, nil
		}
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

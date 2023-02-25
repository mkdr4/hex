package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/mkdr4/hex/internal/domain"
	"github.com/mkdr4/hex/pkg/database"
	"github.com/spf13/viper"
)

func (c *middlewareController) storeAuth(ctx *fiber.Ctx, s string) (bool, error) {

	if !viper.GetBool("httpserver.auth") {
		return true, nil
	}

	authKey, err := database.GetAuthkey(c.db, domain.Medium)
	if err != nil {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}

	for _, key := range authKey {
		if key.Key == s {
			return true, nil
		}
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

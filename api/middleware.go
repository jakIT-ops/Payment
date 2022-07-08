package api

import (
	"errors"
	"fmt"
	"strings"

	"payment_full/token"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.GetReqHeaders()[authorizationHeaderKey]

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		accessToken := fields[1]
		_, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		// Set sets the response's HTTP header field to the specified key, value.
		ctx.Next()
		return nil
	}
}

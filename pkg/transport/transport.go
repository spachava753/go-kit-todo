package transport

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/gofiber/fiber/v2"
)

// MakeFiberHandler creates a fiber handler
func MakeFiberHandler(
	endpoint endpoint.Endpoint,
	decode func(ctx *fiber.Ctx) (interface{}, error),
	encode func(ctx *fiber.Ctx, resp interface{}) error,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		request, err := decode(ctx)
		if err != nil {
			return err
		}

		resp, err := endpoint(nil, request)
		if err != nil {
			return err
		}

		if err := encode(ctx, resp); err != nil {
			return err
		}
		return nil
	}
}

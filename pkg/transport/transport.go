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
	errHandler func(ctx *fiber.Ctx, err error) error,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if errHandler == nil {
			errHandler = func(ctx *fiber.Ctx, err error) error {
				ctx.Status(500)
				return err
			}
		}
		request, err := decode(ctx)
		if err != nil {
			return errHandler(ctx, err)
		}

		resp, err := endpoint(nil, request)
		if err != nil {
			return errHandler(ctx, err)
		}

		if err := encode(ctx, resp); err != nil {
			return errHandler(ctx, err)
		}
		return nil
	}
}

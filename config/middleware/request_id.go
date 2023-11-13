package middleware

import (
	context_lib "github.com/Lyalyashechka/actresses/internal/context"
	"github.com/labstack/echo/v4"
)

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := context.Request().Context()
			requestID := context_lib.GetOrGenerateRequestID(ctx)
			ctx = context_lib.AddCtxFields(ctx, context_lib.Fields{
				"request_id": requestID,
			})

			context.SetRequest(context.Request().WithContext(context_lib.SetRequestID(ctx, requestID)))

			return next(context)
		}
	}
}

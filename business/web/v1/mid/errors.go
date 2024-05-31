package mid

import (
	"context"
	"net/http"

	"github.com/raton1180/service/app/foundation/logger"
	"github.com/raton1180/service/business/web/v1/response"
	"github.com/raton1180/service/foundation/web"
)

// Errors handles errors coming out of the call chain.
func Errors(log *logger.Logger) web.Middleware {
	// Logger executes the logger middleware functionality.
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {
				log.Error(ctx, "message", "msg", err)

				var er response.ErrorDocument
				var status int

				switch {
				case response.IsError(err):
					reqErr := response.GetError(err)
					er = response.ErrorDocument{
						Error: reqErr.Error(),
					}
					status = reqErr.Status
				default:
					er = response.ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Respond(ctx, w, er, status); err != nil {
					return err
				}

				if web.IsShutdown(err) {
					return err
				}
			}
			return nil
		}
		return h
	}
	return m
}

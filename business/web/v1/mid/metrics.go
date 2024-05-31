package mid

import (
	"context"
	"net/http"

	"github.com/raton1180/service/business/web/v1/metrics"
	"github.com/raton1180/service/foundation/web"
)

// Metrics updates program counters.
func Metrics() web.Middleware {
	// Logger executes the logger middleware functionality.
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx = metrics.Set(ctx)

			err := handler(ctx, w, r)

			n := metrics.AddRequest(ctx)

			if n%1000 == 0 {
				metrics.AddErrors(ctx)
			}
			return err
		}
		return h
	}
	return m
}

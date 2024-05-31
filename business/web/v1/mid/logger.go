package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/raton1180/service/app/foundation/logger"
	"github.com/raton1180/service/foundation/web"
)

func Logger(log *logger.Logger) web.Middleware {

	// Logger executes the logger middleware functionality.
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			v := web.GetValues(ctx)

			path := r.URL.Path

			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Info(ctx, "request complete", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))

			return err
		}

		return h
	}
	return m
}

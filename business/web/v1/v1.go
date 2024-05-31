package v1

import (
	"os"

	"github.com/raton1180/service/app/foundation/logger"
	"github.com/raton1180/service/business/web/v1/mid"
	"github.com/raton1180/service/foundation/web"
)

type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

func APIMux(cfg APIMuxConfig, routerAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())
	routerAdder.Add(app, cfg)
	return app
}

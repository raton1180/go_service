package v1

import (
	"os"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/raton1180/service/app/foundation/logger"
)

type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *httptreemux.ContextMux, cfg APIMuxConfig)
}

func APIMux(cfg APIMuxConfig, routerAdder RouteAdder) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()
	routerAdder.Add(mux, cfg)
	return mux
}

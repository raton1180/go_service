package hackgrp

import (
	"net/http"

	"github.com/raton1180/service/foundation/web"
)

func Routes(app *web.App) {
	app.Handle(http.MethodGet, "/hack", Hack)
}

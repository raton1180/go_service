package handlers

import (
	"github.com/raton1180/service/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/raton1180/service/business/web/v1"
	"github.com/raton1180/service/foundation/web"
)

type Routes struct{}

func (Routes) Add(app *web.App, cfg v1.APIMuxConfig) {
	hackgrp.Routes(app)
}

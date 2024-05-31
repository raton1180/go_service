package hackgrp

import (
	"context"
	"net/http"

	"github.com/raton1180/service/foundation/web"
)

func Want(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

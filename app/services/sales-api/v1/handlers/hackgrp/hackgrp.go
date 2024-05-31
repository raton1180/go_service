package hackgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/raton1180/service/business/web/v1/response"
	"github.com/raton1180/service/foundation/web"
)

func Want(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100) % 2; n == 0 {
		return response.NewError(errors.New("TRUST ERROR"), http.StatusBadRequest)
	}
	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

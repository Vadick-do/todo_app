package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/Vadick-do/todo_app/internal/core/transport/htpp/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}

package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Vadick-do/todo_app/internal/core/domain"
	core_http_server "github.com/Vadick-do/todo_app/internal/core/transport/htpp/server"
)

type UsersHTTPHandler struct {
	userService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersHTTPHandler(
	userService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: userService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}

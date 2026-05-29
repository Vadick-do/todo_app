package users_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
)

type GetUserResponse UserDTOResponce

// GetUser       godoc
// @Summary      Получение пользователя
// @Description  Получение конкретного пользователя по его ID
// @Tags         users
// @Produce      json
// @Param        id path int true "ID получаемого пользователя"
// @Success      200 {object} GetUserResponse "Пользователь успешно найден"
// @Failure      400 {object} core_http_responce.ErrorResponse "Bad request"
// @Failure      404 {object} core_http_responce.ErrorResponse "User not found"
// @Failure      500 {object} core_http_responce.ErrorResponse "Internal server error"
// @Router       /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_responce.NewHTTPResponce(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}

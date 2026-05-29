package users_transport_http

import (
	"net/http"

	"github.com/Vadick-do/todo_app/internal/core/domain"
	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"                  example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+" example:"+79998887766"`
}

type CreateUserResponse UserDTOResponce

// CreateUser    godoc
// @Summary      Создать пользователя
// @Description  Создать нового пользователя в системе
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true "CreateUser тело запроса"
// @Success      201 {object} CreateUserResponse "Успешно созданный пользователь"
// @Failure      400 {object} core_http_responce.ErrorResponse "Bad request"
// @Failure      500 {object} core_http_responce.ErrorResponse "Internal server error"
// @Router       /users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_responce.NewHTTPResponce(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responceHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responceHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))
	responceHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitialized(dto.FullName, dto.PhoneNumber)
}

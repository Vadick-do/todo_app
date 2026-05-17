package tasks_transport_http

import (
	"net/http"

	"github.com/Vadick-do/todo_app/internal/core/domain"
	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_responce.NewHTTPResponce(log, rw)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskDomin := domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	taskDomin, err := h.tasksService.CreateTask(ctx, taskDomin)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomin))
	responseHandler.JSONResponse(response, http.StatusOK)
}

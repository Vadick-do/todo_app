package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
)

type GetTaskResponse TaskDTOResponse

// CreateTask   godoc
// @Summary     Получение задачи
// @Description Получение конкретной задачи по ID
// @Tags        tasks
// @Produce     json
// @Param       id path int true "ID получаемой задачи"
// @Success     200 {object} GetTaskResponse "Задача успешно найдена"
// @Failure     400 {object} core_http_responce.ErrorResponse "Bad request"
// @Failure     404 {object} core_http_responce.ErrorResponse "Author not found"
// @Failure     500 {object} core_http_responce.ErrorResponse "Internal server error"
// @Router      /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_responce.NewHTTPResponce(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task id path value")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

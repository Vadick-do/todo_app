package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
)

// CreateTask   godoc
// @Summary     Удаление задачи
// @Description Удалить существующую в системе задачу по ее ID
// @Tags        tasks
// @Param       id path int true "ID удаляемой задачи"
// @Success     204 "Успешное удаление задачи"
// @Failure     400 {object} core_http_responce.ErrorResponse "Bad request"
// @Failure     404 {object} core_http_responce.ErrorResponse "Task not found"
// @Failure     500 {object} core_http_responce.ErrorResponse "Internal server error"
// @Router      /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	reponseHanlder := core_http_responce.NewHTTPResponce(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		reponseHanlder.ErrorResponse(err, "failed to get task id path value")
		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		reponseHanlder.ErrorResponse(err, "failed to delete task")
		return
	}

	reponseHanlder.NoContentResponse()
}

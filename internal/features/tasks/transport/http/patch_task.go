package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/Vadick-do/todo_app/internal/core/domain"
	core_logger "github.com/Vadick-do/todo_app/internal/core/logger"
	core_http_request "github.com/Vadick-do/todo_app/internal/core/transport/htpp/request"
	core_http_responce "github.com/Vadick-do/todo_app/internal/core/transport/htpp/responce"
	core_http_types "github.com/Vadick-do/todo_app/internal/core/transport/htpp/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"            swaggertype:"string"  example:"Погулять с собакой"`
	Description core_http_types.Nullable[string] `json:"description"      swaggertype:"string"  example:"null"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"        swaggertype:"boolean"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title can not be null")
		}

		titileLen := len([]rune(*r.Title.Value))
		if titileLen < 1 || titileLen > 100 {
			return fmt.Errorf("`Title` must be between 1 and 100 symbols")
		}

	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("`Description` must be between 1 and 1000 symbols")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("`Completed` can not be null")
		}
	}

	return nil
}

type PatchUserResponse TaskDTOResponse

// PatchTask    godoc
// @Summary     Обновить задачу
// @Description Обновляет информацию об уже существующей в системе задаче
// @Description ### Логика обновления полей (Three-state-logic):
// @Description 1. **Поле не передано**: `description` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `"description": "Утром в 06:30 выйти на прогулку с Шариком"` - устанавливает новое описание для задачи
// @Description 3. **Явно передан null**: `"description": null` - очищает поле в БД (set to NULL)
// @Description Ограничения: `title` и `completed` не могут быть выставлены как null
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param       id path int true "ID изменяемой задачи"
// @Param       request body PatchTaskRequest true "PatchTask тело запроса"
// @Success     200 {object} PatchUserResponse "Успешно измененная задача"
// @Failure     400 {object} core_http_responce.ErrorResponse "Bad request"
// @Failure     404 {object} core_http_responce.ErrorResponse "Task not found"
// @Failure     500 {object} core_http_responce.ErrorResponse "Internal server error"
// @Router      /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_responce.NewHTTPResponce(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task id path value")
		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}

	response := PatchUserResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}

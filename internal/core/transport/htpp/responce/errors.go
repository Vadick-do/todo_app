package core_http_responce

type ErrorResponse struct {
	Error   string `json:"error"   example:"full error text"`
	Message string `json:"message" example:"short human-readable message"`
}

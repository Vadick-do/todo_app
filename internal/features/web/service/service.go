package web_service

type WebServie struct {
	webRepository WebRepository
}

type WebRepository interface {
	GetFile(filePath string) ([]byte, error)
}

func NewWebService(
	webRepository WebRepository,
) *WebServie {
	return &WebServie{
		webRepository: webRepository,
	}
}

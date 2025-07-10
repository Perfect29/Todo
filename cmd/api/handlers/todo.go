package handlers

import(
	"github.com/Perfect29/Server/cmd/api/service"
)

type Handler struct {
	todoUsecase *service.TodoUsecase
}

func NewHandler(todoUsecase *service.TodoUsecase) *Handler {
	return &Handler{
		todoUsecase: todoUsecase,
	}
}
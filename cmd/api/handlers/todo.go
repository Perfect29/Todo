package handlers

import(
	"github.com/Perfect29/Server/cmd/api/service"
)

type Handler struct {
	Srv *service.TodoService
}
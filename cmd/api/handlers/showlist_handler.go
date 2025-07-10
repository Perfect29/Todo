package handlers

import (
	"github.com/labstack/echo"
	
	"net/http"
)

func (h *Handler) ShowlistHandler(c echo.Context) error {
	todos, err := h.todoUsecase.ShowListTodo(c.Request().Context())

	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to get List of Todos"
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, todos)
}
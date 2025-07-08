package handlers

import (
	"github.com/labstack/echo"

	"net/http"
)

func (h *Handler) ShowlistHandler(c echo.Context) error {
	todos, err := h.Srv.ShowListTodo()

	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to get List of Todos"
	}

	return c.JSON(http.StatusOK, todos)
}
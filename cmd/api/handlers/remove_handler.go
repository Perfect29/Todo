package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) RemoveHandler(c echo.Context) error {
	id := c.Param("id")
	err := h.Srv.RemoveTodo(id)
	res := make(map[string]any)
	res["message"] = "todo deleted"
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"error": "todo was not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}
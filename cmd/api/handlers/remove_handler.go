package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) RemoveHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		res := make(map[string]any)
		res["error"] = "Invalid id parameter"
		return c.JSON(http.StatusBadRequest, res)
	}
	err = h.Service.RemoveTodo(context.Background(), idx)
	res := make(map[string]any)
	res["message"] = "todo deleted"
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"error": "todo was not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}
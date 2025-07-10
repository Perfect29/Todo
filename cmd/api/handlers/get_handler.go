package handlers

import (
	"context"
	"strconv"

	"github.com/labstack/echo"

	"net/http"
)

func (h *Handler) GetHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		res := make(map[string]any)
		res["error"] = "Bad Request"
		return c.JSON(http.StatusBadRequest, res)
	}
	todo, err := h.todoUsecase.GetByID(context.Background(), idx)

	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to get a todo by id"
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, todo)
}
package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) RemoveHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf("Invalid id parameter in RemoveHandler: %s, %v", id, err)
		res := make(map[string]any)
		res["error"] = "Invalid id parameter"
		return c.JSON(http.StatusBadRequest, res)
	}
	err = h.todoUsecase.RemoveTodo(c.Request().Context(), idx)
	res := make(map[string]any)
	res["message"] = "todo deleted"
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"error": "todo was not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}
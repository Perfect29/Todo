package handlers

import (
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func (h *Handler) GetHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf("Invalid id parameter in GetHandler: %s, %v", id, err)
		res := make(map[string]any)
		res["error"] = "Bad Request"
		return c.JSON(http.StatusBadRequest, res)
	}
	todo, err := h.todoUsecase.GetByID(c.Request().Context(), idx)

	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to get a todo by id"
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, todo)
}
package handlers

import (
	"net/http"

	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/labstack/echo"
)

func (h *Handler) AddHandler(c echo.Context) error {
	var todo service.Todo 
	if err := c.Bind(&todo); err != nil {
		res := make(map[string]string)
		res["error"] = "Invalid Request"
		return c.JSON(http.StatusBadRequest, res)
	}

	err :=  h.Service.AddTodo(c.Request().Context(), &todo)
	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to add todo"
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, todo)
}
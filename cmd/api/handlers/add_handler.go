package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/Perfect29/Server/cmd/api/service"
)

func (h *Handler) AddHandler(c echo.Context) error {
	var todo service.Todo 
	if err := c.Bind(&todo); err != nil {
		res := make(map[string]string)
		res["error"] = "Invalid Request"
		return c.JSON(http.StatusBadRequest, res)
	}

	err :=  h.Srv.AddTodo(&todo)
	if err != nil {
		res := make(map[string]string)
		res["error"] = "Failed to add todo"
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, todo)
}
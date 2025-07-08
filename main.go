package main

import (
	// "context"
	// "log"

	"github.com/Perfect29/Server/cmd/api/handlers"
	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	// db, err := service.InitDB(context.Background())
	// if err != nil {
	// 	log.Fatal("Unable to connect to database:", err)
	// }
	// defer db.Close(context.Background())
	
	todoService := &service.TodoService{}
	h := &handlers.Handler{
		Srv: todoService,
	}

	e := echo.New()
	e.POST("/add",h.AddHandler)
	e.DELETE("/remove/:id", h.RemoveHandler)
	e.GET("/showlist", h.ShowlistHandler)

	e.Logger.Fatal(e.Start(":1323"))

}

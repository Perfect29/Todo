package main

import (
	"context"
	"log"

	"github.com/Perfect29/Server/cmd/api/handlers"
	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	dbService, err := service.InitDB(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer dbService.Close(ctx)

	cacheService := service.NewCacheService()

	todoService := &service.Service{
		DB:    dbService.DB, 
		Cache: cacheService,
	}
	h := &handlers.Handler{
		Service: todoService,
	}

	e := echo.New()
	e.POST("/add",h.AddHandler)
	e.DELETE("/remove/:id", h.RemoveHandler)
	e.GET("/showlist", h.ShowlistHandler)
	e.GET("/get/:id", h.GetHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

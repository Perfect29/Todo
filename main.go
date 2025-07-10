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


	postgresRepo, err := service.InitDB(ctx)
	if err != nil {
		log.Fatal("Could not connect to Database", err)
	}
	cacheRepo := service.NewCacheRepository()

	todoUsecase := service.TodoUsecase{
		Repo: postgresRepo,
		Cache: cacheRepo,
	}
	h := handlers.NewHandler(&todoUsecase)

	e := echo.New()
	e.POST("/add",h.AddHandler)
	e.DELETE("/remove/:id", h.RemoveHandler)
	e.GET("/showlist", h.ShowlistHandler)
	e.GET("/get/:id", h.GetHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

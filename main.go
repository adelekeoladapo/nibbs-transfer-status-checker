package main

import (
	"ajocard/status-indicator/controller"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	/* Load configuration */
	err:= godotenv.Load()
	if err != nil {
		log.Panic("Could not load app configuration")
	}

	/* App Router */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello AjoCard")
	})
	e.GET("/api/v1/transfer-status", func(context echo.Context) error {
		return controller.CheckTransferStatus(context)
	})

	e.Logger.Fatal(e.Start(":8090"))
	
}

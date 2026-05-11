package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// โหลด .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// ดึง PORT จาก .env
	port := os.Getenv("PORT")

	e := echo.New()

	// Health Check API
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.Logger.Fatal(e.Start(":" + port))
}

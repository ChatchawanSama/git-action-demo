package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// โหลด .env ถ้ามี
	_ = godotenv.Load()

	// ดึง PORT จาก ENV
	port := os.Getenv("PORT")

	// default port
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	// Health Check API
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.Logger.Fatal(e.Start(":" + port))
}

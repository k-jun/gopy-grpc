package main

import (
	"log"
	"os"

	"gopy-grpc-server/infrastructure/grpc"
	"gopy-grpc-server/infrastructure/routers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	done, err := grpc.Initialize()
	defer done()
	if err != nil {
		log.Fatalf("failed to access to grpc-server: %v", err)
	}

	routers.IndexRouting(e)

	port := "8080"
	if os.Getenv("GO_PORT") != "" {
		port = os.Getenv("GO_PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}

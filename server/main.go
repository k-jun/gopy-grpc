package main

import (
	"log"

	"github.com/K-jun1221/ca-adtech-comp/server/infrastructure/grpc"
	"github.com/K-jun1221/ca-adtech-comp/server/infrastructure/routers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	done, err := grpc.Initialize("127.0.0.1:50051")
	defer done()
	if err != nil {
		log.Fatalf("failed to access to grpc-server: %v", err)
	}

	routers.IndexRouting(e)
	routers.PredictRouting(e.Group("/predict"))

	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"log"
	"os"

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

	grpcHost := "127.0.0.1"
	grpcPort := "50051"
	if os.Getenv("GRPC_HOST") != "" {
		grpcHost = os.Getenv("GRPC_HOST")
	}
	if os.Getenv("GRPC_PORT") != "" {
		grpcPort = os.Getenv("GRPC_PORT")
	}

	grpcHost2 := "127.0.0.1"
	grpcPort2 := "50052"
	if os.Getenv("GRPC_HOST2") != "" {
		grpcHost2 = os.Getenv("GRPC_HOST2")
	}
	if os.Getenv("GRPC_PORT2") != "" {
		grpcPort2 = os.Getenv("GRPC_PORT2")
	}
	
	done, done2, err := grpc.Initialize(grpcHost+":"+grpcPort, grpcHost2+":"+grpcPort2)
	defer done()
	defer done2()
	if err != nil {
		log.Fatalf("failed to access to grpc-server: %v", err)
	}

	routers.IndexRouting(e)
	routers.PredictRouting(e.Group("/predict"))

	port := "8080"
	if os.Getenv("GO_PORT") != "" {
		port = os.Getenv("GO_PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}

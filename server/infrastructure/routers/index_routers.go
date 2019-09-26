package routers

import (
	"gopy-grpc-server/interface/controllers"

	"github.com/labstack/echo"
)

func IndexRouting(e *echo.Echo) {
	e.GET("/", controllers.HelloEcho)
	e.GET("/predict", controllers.PredictIris)
}

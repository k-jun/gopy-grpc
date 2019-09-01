package routers

import (
	"github.com/K-jun1221/ca-adtech-comp/server/interface/controllers"

	"github.com/labstack/echo"
)

func PredictRouting(e *echo.Group) {
	e.GET("", controllers.PredictIris)
}

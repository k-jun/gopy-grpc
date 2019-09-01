package controllers

import (
	"net/http"

	"github.com/K-jun1221/ca-adtech-comp/server/infrastructure/grpc"
	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"

	"github.com/labstack/echo"
)

func PredictIris(c echo.Context) error {
	arg := &adtech.Request{
		SepalLength: 1.0,  // not 0.0
		SepalWidth:  1.0,  // not 0.0
		PetalLength: 0.5,  // not 0.0
		PetalWidth:  10.0, // not 0.0
	}
	res, err := grpc.Predict(arg)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

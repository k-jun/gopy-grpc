package controllers

import (
	"net/http"

	"github.com/K-jun1221/ca-adtech-comp/server/infrastructure/grpc"
	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"

	"github.com/labstack/echo"
)

func PredictIris(c echo.Context) error {
	arg := &adtech.Request{
		SepalLength: 4.8, // not 0.0
		SepalWidth:  3.0, // not 0.0
		PetalLength: 1.4, // not 0.0
		PetalWidth:  0.1, // not 0.0
	}
	res, err := grpc.Predict(arg)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

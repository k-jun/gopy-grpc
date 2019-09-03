package controllers

import (
	"net/http"

	"github.com/K-jun1221/ca-adtech-comp/server/infrastructure/grpc"
	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"

	"github.com/labstack/echo"
)

func PredictIris(c echo.Context) error {
	// 6.7 2.5 5.8 1.8 -> Virginica
	arg := &adtech.Request{
		SepalLength: 6.7, // not 0.0
		SepalWidth:  2.5, // not 0.0
		PetalLength: 5.8, // not 0.0
		PetalWidth:  1.8, // not 0.0
	}
	res, err := grpc.Predict(arg)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

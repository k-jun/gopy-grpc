package controllers

import (
	"net/http"

	"gopy-grpc-server/infrastructure/grpc"
	proto "gopy-grpc-server/protolib"

	"github.com/labstack/echo"
)

type simpleResponse struct {
	Message string `json: "message"`
}

func HelloEcho(c echo.Context) error {
	return c.JSON(http.StatusOK, simpleResponse{Message: "this api is working! Ver 1.0.0"})
}

func PredictIris(c echo.Context) error {
	// 6.7 2.5 5.8 1.8 -> Virginica
	arg := &proto.Request{
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

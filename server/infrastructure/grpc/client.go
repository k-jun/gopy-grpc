package grpc

import (
	"context"
	"fmt"
	"time"

	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"

	"google.golang.org/grpc"
)

var Client adtech.AdTechClient

func Initialize(serverAddr string) (func() error, error) {
	var opts []grpc.DialOption
	// TODO TLS試してみる...? Connection uses plain TCP, TLS also exists
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)

	fmt.Println(conn, err)
	if err != nil {
		return nil, err
	}

	Client = adtech.NewAdTechClient(conn)
	return conn.Close, nil
}

func Predict(params *adtech.Request) (*adtech.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := Client.Predict(ctx, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}

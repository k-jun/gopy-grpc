package grpc

import (
	"context"
	"time"

	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"

	"google.golang.org/grpc"
)

var Client adtech.AdTechClient

func Initialize(serverAddr string) (func() error, error) {
	var opts []grpc.DialOption
	// TODO TLS試してみる...? Connection uses plain TCP, TLS also exists
	// TODO ClientSideLBやる 参考(https://deeeet.com/writing/2018/03/30/kubernetes-grpc/)
	// TODO GoRoutineを使って並列に複数のモデルの結果を集める。
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
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

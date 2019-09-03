package grpc

import (
	"context"
	"time"

	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"
	"google.golang.org/grpc"
)

var clientSvm adtech.AdTechClient

func initSvm(serverAddr string) (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	clientSvm = adtech.NewAdTechClient(conn)
	return conn.Close, nil
}

func predictSvm(params *adtech.Request, ch chan grpcChan) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := clientSvm.Predict(ctx, params)
	ch <- grpcChan{Response: response, Err: err}
}

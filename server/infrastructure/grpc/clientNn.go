package grpc

import (
	"context"
	"time"

	proto "gopy-grpc-server/protolib"

	"google.golang.org/grpc"
)

var clientNn proto.ProtoClient

func initNn(serverAddr string) (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	clientNn = proto.NewProtoClient(conn)
	return conn.Close, nil
}

func predictNn(params *proto.Request, ch chan grpcChan) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := clientNn.Predict(ctx, params)
	ch <- grpcChan{Response: response, Err: err}
}

package grpc

import (
	"context"
	"time"

	proto "gopy-grpc-server/protolib"

	"google.golang.org/grpc"
)

var clientSvm proto.ProtoClient

func initSvm(serverAddr string) (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	clientSvm = proto.NewProtoClient(conn)
	return conn.Close, nil
}

func predictSvm(params *proto.Request, ch chan grpcChan) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := clientSvm.Predict(ctx, params)
	ch <- grpcChan{Response: response, Err: err}
}

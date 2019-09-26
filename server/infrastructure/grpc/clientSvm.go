package grpc

import (
	"context"
	"time"

	proto "gopy-grpc-server/protolib"

	"google.golang.org/grpc"
)

var clientSvm proto.ProtoClient

type SVM struct {
	Conn *grpc.ClientConn
}

func newSvm(serverAddr string) (func() error, ML, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	ml := &SVM{Conn: conn}
	if err != nil {
		return nil, ml, err
	}
	clientSvm = proto.NewProtoClient(conn)
	return conn.Close, ml, nil
}

func (ml *SVM) Predict(params *proto.Request) (*proto.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := clientSvm.Predict(ctx, params)
	return response, err
}

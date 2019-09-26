package grpc

import (
	"fmt"
	"gopy-grpc-server/common"
	proto "gopy-grpc-server/protolib"
)

type grpcChan struct {
	Response *proto.Response
	Err      error
}

type ML interface {
	Predict(params *proto.Request) (*proto.Response, error)
}

var ml1 ML
var ml2 ML

func Initialize() (func() error, error) {
	grpcHost := common.GetEnv("GRPC_HOST", "127.0.0.1")
	grpcPort := common.GetEnv("GRPC_PORT", "50051")

	grpcHost2 := common.GetEnv("GRPC_HOST2", "127.0.0.1")
	grpcPort2 := common.GetEnv("GRPC_PORT2", "50052")

	doneSvm1, ml1, err := newSvm(grpcHost + ":" + grpcPort)
	doneSvm2, ml2, err := newSvm(grpcHost2 + ":" + grpcPort2)
	if err != nil {
		return nil, err
	}

	fmt.Println(ml1, ml2)

	done := func() error {
		err1 := doneSvm1()
		err2 := doneSvm2()

		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
		return nil
	}

	return done, nil
}

func Predict(params *proto.Request) (*proto.Response, error) {

	// chNn := make(chan grpcChan)
	// chSvm := make(chan grpcChan)

	// go predictNn(params, chNn)
	// go predictSvm(params, chSvm)

	// resultNn := <-chNn
	// resultSvm := <-chSvm
	// close(chNn)
	// close(chSvm)

	// // TODO２種類のエラーをうまく扱う
	// itisType := ""
	// if resultNn.Err != nil && resultSvm.Err != nil {
	// 	return nil, resultNn.Err
	// }
	// if resultNn.Err != nil {
	// 	itisType += resultNn.Err.Error()
	// } else {
	// 	itisType += resultNn.Response.IrisType
	// }
	// if resultSvm.Err != nil {
	// 	itisType += resultSvm.Err.Error()
	// } else {
	// 	itisType += resultSvm.Response.IrisType
	// }

	return &proto.Response{IrisType: "itisType"}, nil
}

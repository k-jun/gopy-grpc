package grpc

import (
	adtech "github.com/K-jun1221/ca-adtech-comp/server/protolib"
)

type grpcChan struct {
	Response *adtech.Response
	Err      error
}

func Initialize(serverAddr string, serverAddr2 string) (func() error, func() error, error) {
	// TODO TLS試してみる...? Connection uses plain TCP, TLS also exists
	// TODO ClientSideLBやる 参考(https://deeeet.com/writing/2018/03/30/kubernetes-grpc/)
	doneNn, err := initNn(serverAddr)
	doneSvm, err := initSvm(serverAddr2)
	if err != nil {
		return nil, nil, err
	}

	return doneNn, doneSvm, nil
}

func Predict(params *adtech.Request) (*adtech.Response, error) {

	chNn := make(chan grpcChan)
	chSvm := make(chan grpcChan)

	go predictNn(params, chNn)
	go predictSvm(params, chSvm)

	resultNn := <-chNn
	resultSvm := <-chSvm
	close(chNn)
	close(chSvm)

	// TODO２種類のエラーをうまく扱う
	itisType := ""
	if resultNn.Err != nil && resultSvm.Err != nil {
		return nil, resultNn.Err
	}
	if resultNn.Err != nil {
		itisType += resultNn.Err.Error()
	} else {
		itisType += resultNn.Response.IrisType
	}
	if resultSvm.Err != nil {
		itisType += resultSvm.Err.Error()
	} else {
		itisType += resultSvm.Response.IrisType
	}

	return &adtech.Response{IrisType: itisType}, nil
}

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
	// opts = append(opts, grpc.WithInsecure())
	doneNn, err := initNn(serverAddr)
	doneSvm, err := initSvm(serverAddr)
	if err != nil {
		return nil, nil, err
	}

	// Client = adtech.NewAdTechClient(conn)
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
	if resultNn.Err != nil {
		return nil, resultNn.Err
	}
	if resultSvm.Err != nil {
		return nil, resultSvm.Err
	}

	return &adtech.Response{IrisType: resultNn.Response.IrisType + " / " + resultSvm.Response.IrisType}, nil
}

package grpc

import (
	"gopy-grpc-server/common"
	proto "gopy-grpc-server/protolib"
	"strconv"
)

type grpcChan struct {
	Response *proto.Response
	Err      error
}

type ML interface {
	Predict(params *proto.Request) (*proto.Response, error)
}

var mlMulti []ML
var mlSingle ML

func Initialize() (func() error, error) {
	voting := common.GetEnv("GRPC_VOTING", "false")

	if voting == "true" {
		done, err := initVoting()
		return done, err
	}
	done, err := initSingle()
	return done, err
}

func initVoting() (func() error, error) {
	numStr := common.GetEnv("GRPC_VOTING_AMMOUNT", "1")
	num, err := strconv.Atoi(numStr)

	if err != nil {
		return nil, err
	}

	doneFuncs := [](func() error){}

	for i := 1; i <= num; i++ {
		grpcHost := common.GetEnv("GRPC_HOST_"+strconv.Itoa(i), "127.0.0.1")
		grpcPort := common.GetEnv("GRPC_PORT_"+strconv.Itoa(i), "50051")
		grpcType := common.GetEnv("GRPC_TYPE_"+strconv.Itoa(i), "svm")
		done, ml, err := selectModel(grpcHost+":"+grpcPort, grpcType)
		if err != nil {
			return nil, err
		}
		mlMulti = append(mlMulti, ml)
		doneFuncs = append(doneFuncs, done)

	}

	done := func() error {
		for _, f := range doneFuncs {
			err := f()
			if err != nil {
				return err
			}
		}
		return nil
	}

	return done, nil
}

func initSingle() (func() error, error) {
	grpcHost := common.GetEnv("GRPC_HOST_SINGLE", "127.0.0.1")
	grpcPort := common.GetEnv("GRPC_PORT_SINGLE", "50051")
	grpcType := common.GetEnv("GRPC_TYPE_SINGLE", "svm")
	done, ml, err := selectModel(grpcHost+":"+grpcPort, grpcType)
	mlSingle = ml
	if err != nil {
		return nil, err
	}

	return done, nil

}

func selectModel(address string, mlType string) (func() error, ML, error) {
	if mlType == "svm" {
		return newSvm(address)
	}
	// default
	return newSvm(address)
}

func Predict(params *proto.Request) (*proto.Response, error) {
	voting := common.GetEnv("GRPC_VOTING", "false")

	if voting != "true" {
		return mlSingle.Predict(params)
	}

	results := make(chan grpcChan)

	for _, ml := range mlMulti {
		go func() {
			req, err := ml.Predict(params)
			results <- grpcChan{Response: req, Err: err}
		}()
	}

	var votingResult string
	for i := 1; i <= len(mlMulti); i++ {
		result := <-results
		if result.Err != nil {
			return &proto.Response{IrisType: ""}, result.Err
		}
		votingResult = votingResult + result.Response.IrisType
	}
	close(results)

	return &proto.Response{IrisType: votingResult}, nil
}

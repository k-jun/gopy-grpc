# golang-server-template

```sh
# kubernetesのversionを確認して起動
gcloud container get-server-config --zone asia-northeast1-a
gcloud container clusters create k8s --cluster-version 1.13.7-gke.8 --zone asia-northeast1-a --num-nodes 3

# kubernetesのCredentialを~/.kube/configに保存
gcloud container clusters get-credentials k8s --zone asia-northeast1-a
```


```sh
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

pip install grpcio-tools
```

### `adtech.pb.go`を生成
```
protoc -I protos/ --go_out=plugins=grpc:./server/protolib protos/adtech.proto
```

### `adtech_pb2_grpc.py`と`adtech_pb2.py`を生成
```sh
python -m grpc_tools.protoc -I./protos --python_out=./ml/protolib --grpc_python_out=./ml/protolib ./protos/adtech.proto
```

Module化に伴い`XXXX_pb2_grpc`のimportを適宜修正

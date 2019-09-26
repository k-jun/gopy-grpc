# gopy-grpc-server

## **Download**

```sh
git clone https://github.com/K-jun1221/gopy-grpc-server.git
cd gopy-grpc-server
```

## **ProtocolBuffer**

```sh
# golang
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

# `proto.pb.go`を生成
mkdir ./server/protolib
protoc -I protos/ --go_out=plugins=grpc:./server/protolib protos/proto.proto

# python
python -m pip install --upgrade pip
python -m pip install grpcio
python -m pip install grpcio-tools

# `proto_pb2.py`と`proto_pb2_grpc.py`を生成
python -m grpc_tools.protoc -I./protos --python_out=./ml/protolib --grpc_python_out=./ml/protolib ./protos/proto.proto

```

#### ※Pythonは注意が必要

Module化に伴い`/ml/protolib/proto_pb2_grpc.py`のimportを適宜修正

**変更前**
```python
import proto_pb2 as proto__pb2
```
**変更後**
```python
import protolib.proto_pb2 as proto__pb2
```

- Python参考: https://grpc.io/docs/quickstart/python/
- Golang参考: https://grpc.io/docs/quickstart/go/


## **Activate**

環境変数は`./env.sh`を適宜いじる。

```sh
# golang
cd server
source ../env.sh && go run ./main.go

# python
cd ml
source ../env.sh && python3 ./main.py
```

<!-- TODO 修正する -->
### Deploy

```sh
# kubernetesのversionを確認して起動
gcloud container get-server-config --zone asia-northeast1-a
gcloud container clusters create k8s --cluster-version 1.13.7-gke.8 --zone asia-northeast1-a --num-nodes 3

# kubernetesのCredentialを~/.kube/configに保存
gcloud container clusters get-credentials k8s --zone asia-northeast1-a
```

```
kubectl apply -f kubernetes/deploy.yaml --prune --all # update, create, or delete
```


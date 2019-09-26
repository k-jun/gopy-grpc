# gopy-grpc-server

### SETUP

```sh
https://github.com/K-jun1221/ca-adtech-comp.git
cd ca-adtech-comp
```

gRPCを使っているのでまず、gRPCが入っていないと話に
ならない。実行テストがしたいので、できればPythonの方だけでも入れて欲しい。

```sh
# golang
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

# python
python -m pip install --upgrade pip
python -m pip install grpcio
python -m pip install grpcio-tools

```


Python参考: https://grpc.io/docs/quickstart/python/

Golang参考: https://grpc.io/docs/quickstart/go/

### ProtoLibを作成
protoファイルにしたがって自動生成されるファイルはGitIgnoreしてある。以下のコマンドでそれぞれのgRPC用ライブラリを作成する。

```sh
# `adtech.pb.go`を生成
mkdir ./server/protolib
protoc -I protos/ --go_out=plugins=grpc:./server/protolib protos/adtech.proto
```
```sh
# `adtech_pb2.py`と`adtech_pb2_grpc.py`を生成
python -m grpc_tools.protoc -I./protos --python_out=./ml/protolib --grpc_python_out=./ml/protolib ./protos/adtech.proto
```

#### ※Pythonは注意が必要

Module化に伴い`/ml/protolib/adtech_pb2_grpc.py`のimportを適宜修正

**変更前**
```python
import adtech_pb2 as adtech__pb2
```
**変更後**
```python
import protolib.adtech_pb2 as adtech__pb2
```



### 起動

環境変数は`/env.sh`を適宜いじる。

```sh
# golang
cd server
source ../env.sh && go run ./main.go

# python
cd ml
source ../env.sh && python3 ./main.py
```

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

### TODO

Should
- 機械学習のコードを用意する。
- Docker化してKubernetes試したい
- GoRoutineを試したい 
- DeployをRolling Upgrade -> GreenBlueに変更する。 参考(https://deeeet.com/writing/2018/03/30/kubernetes-grpc/)
- 

Want
- UNIXDomainソケットってどうなの？
- TCPよりもっと良い接続方法は?
- STLで接続したい
- ログをどうやって出す？
- 負荷をかけて検証したい
- 監視ツールを導入したい


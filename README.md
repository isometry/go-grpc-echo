# go-grpc-echo

Simple Go-based gRPC Echo service, designed for k8s/istio testing

## Usage

```console
kubectl create deployment echo --image=ghcr.io/isometry/go-grpc-echo:latest --port=8080
kubectl create service loadbalancer echo --tcp=8080:8080

grpcurl -plaintext localhost:8080 describe
grpcurl -plaintext -d  '{"message": "Hello, World"}' localhost:8080 rpc.Echo/Request

ghz --insecure --call rpc.Echo/Request --data='' --total=10000 --connections=50 --count-errors localhost:8080
```

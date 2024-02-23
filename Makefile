export KO_DOCKER_REPO ?= ghcr.io/isometry/go-grpc-echo

build: rpc/echo.pb.go rpc/echo_grpc.pb.go
	ko build --bare

rpc/echo.pb.go rpc/echo_grpc.pb.go: rpc/echo.proto
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--proto_path=. \
		rpc/echo.proto

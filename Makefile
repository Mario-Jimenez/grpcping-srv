
.PHONY: proto test


proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/v1/ping/ping.proto

build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/pingserver/pingserver ./cmd/pingserver

run:
	./cmd/pingserver/pingserver

test:
	go test -v ./... -cover

docker-build:
	docker build . --build-arg SSH_PRIVATE_KEY --build-arg GOPRIVATE -t grpcping-srv:0.0.1

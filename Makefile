
.PHONY: proto test


proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/v1/ping/ping.proto

build: proto

	CGO_ENABLED=0 go build -o ./cmd/pingserver/pingserver ./cmd/pingserver

run:
	./cmd/pingserver/pingserver

test:
	go test -v ./... -cover

tag := 0.0.2

docker-build:
	docker build . -t ghcr.io/mario-jimenez/grpcping-srv:$(tag)

docker-push:
	docker push ghcr.io/mario-jimenez/grpcping-srv:$(tag)

docker-run:
	docker run --rm ghcr.io/mario-jimenez/grpcping-srv:$(tag)

FROM golang:1.17.5-buster AS build-env
WORKDIR /pingserver
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/pingserver ./cmd/pingserver

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=build-env /go/bin/pingserver /pingserver
COPY --from=build-env /bin/grpc_health_probe /grpc_health_probe
CMD ["/pingserver"]

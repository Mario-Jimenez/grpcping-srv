FROM golang:1.17.5-buster AS build-env
WORKDIR /pingserver
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/pingserver ./cmd/pingserver

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=build-env /go/bin/pingserver /pingserver
CMD ["/pingserver"]

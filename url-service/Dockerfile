ARG GO_VERSION=1.19.0

FROM golang:${GO_VERSION} AS dev
RUN apt-get update && apt-get install -y protobuf-compiler
WORKDIR /url-service
# install reflex for hot reloading
RUN go install github.com/cespare/reflex@latest
# install grpc code generator
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
COPY /url-service/ .
RUN go mod tidy
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    handlers/grpc/qrgen/qrgen.proto

FROM dev as build
WORKDIR /url-service
RUN CGO_ENABLED=0 GOOS=linux go build \ 
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -o urls .

# FROM gcr.io/distroless/static AS prod
FROM alpine:latest AS prod
WORKDIR /url-service
COPY --from=build /url-service .
CMD ["./urls"]

ARG GO_VERSION=1.19.0

FROM golang:${GO_VERSION} AS dev
RUN apt-get update && apt-get install -y protobuf-compiler

# set working directory
WORKDIR /app

# copy files to container
COPY . .

# install reflex for hot reloading
RUN go install github.com/cespare/reflex@latest

# install grpc code generator
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# download dependencies
RUN go mod tidy

# compile grpc .proto files
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/grpc/qrgen/qrgen.proto

FROM dev as build
RUN CGO_ENABLED=0 GOOS=linux go build \ 
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -o urls .

FROM gcr.io/distroless/static AS prod
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /app/urls .
CMD ["./urls"]

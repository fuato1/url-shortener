FROM golang:1.18.0-alpine AS dev
WORKDIR /app
RUN go install github.com/cespare/reflex@latest
COPY . .

FROM dev as build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /app/cmd/main.go

FROM alpine:latest AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app .
COPY --from=build /app/.env .
CMD ["./app"]

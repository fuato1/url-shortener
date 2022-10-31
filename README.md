# shorturl

shorturl is a URL shortener service written in Go which uses Redis as datastore.

Env variables required to run the service:
- APP_PORT (service port)
- REDIS_HOST (redis host)
- REDIS_PORT (redis port)

Optinal env variables to interact with [qrgen](github.com/juanjoss/qrgen):
- GRPC_SERVER_HOST (qrgen host)
- GRPC_SERVER_PORT (qrgen grpc server port)

Endpoints:
- GET / (main page)
- GET /url (get all shortened URLs)
- POST /url (shorten URL)
- GET /{id} (resolve to original URL and redirect)
- GET /qr/{id} (get QR code, requires [qrgen](github.com/juanjoss/qrgen) GRPC server)

Run the development environment with docker compose:

```bash
docker compose up -d
```

or pull and run the image

```bash
docker pull jujoss/shorturl:latest
```

```bash
docker run -dp 80:8080 \ 
    --env APP_PORT=8080 \ 
    --env REDIS_HOST=redis \ 
    --env REDIS_PORT=6379 \ 
    --env GRPC_SERVER_HOST=qrgen \ 
    --env GRPC_SERVER_PORT=8081 \ 
    jujoss/shorturl:latest
```

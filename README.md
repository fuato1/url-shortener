# shorturl

shorturl is a URL shortener service written in Go which uses Redis as datastore.

env variables required to run the service
- APP_PORT (service port)
- REDIS_HOST (redis IP)
- REDIS_PORT (redis port)

optinal env variables to interact with github.com/juanjoss/qrgen
- GRPC_SERVER_HOST (grpc server ip)
- GRPC_SERVER_PORT (grpc server port)

endpoints
- GET / (main page)
- GET /url (get all shortened URLs)
- POST /url (shorten URL)
- GET /{id} (resolve to original URL and redirect)
- GET /qr/{id} (get QR code, requires github.com/juanjoss/qrgen GRPC server)

run in development environment with docker compose
> `docker compose -f docker-compose.dev.yml up -d`

or the production one...
> `docker compose -f docker-compose.prod.yml up -d`

or build and run the image
> `docker build -t shorturl:latest .`

> `docker run -dp 3000:3000 --env APP_PORT= --env REDIS_HOST= --env REDIS_PORT=`
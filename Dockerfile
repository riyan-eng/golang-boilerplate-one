# build
FROM golang:1.21.1-bullseye AS build-stage
WORKDIR /app
COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -o ./pkg/docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM debian:buster-slim AS release-stage
WORKDIR /
COPY --from=build-stage /binary /binary
COPY --from=build-stage /app/casbin.conf /casbin.conf
COPY --from=build-stage /app/env.json /env.json
CMD ["/binary"]
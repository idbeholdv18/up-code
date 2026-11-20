FROM golang:1.27.3-alpine3.24 AS build

WORKDIR /usr/local/src
COPY shared/go.mod ./shared/
COPY notifications/go.mod notifications/go.sum ./notifications/

WORKDIR /usr/local/src/shared
RUN go mod download

WORKDIR /usr/local/src/notifications
RUN go mod download

WORKDIR /usr/local/src
COPY shared/ ./shared/
COPY notifications/ ./notifications/

WORKDIR /usr/local/src/notifications
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /usr/local/bin/notifications ./cmd/worker/main.go

FROM alpine:3.24

WORKDIR /usr/local
COPY --from=build /usr/local/bin/notifications /usr/local/bin/notifications

EXPOSE 8080

CMD [ "/usr/local/bin/notifications" ]



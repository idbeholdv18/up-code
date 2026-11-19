FROM golang:1.27.3-alpine3.24 AS build

WORKDIR /usr/local/src
COPY shared/go.mod ./shared/
COPY app/go.mod app/go.sum ./app/

WORKDIR /usr/local/src/shared
RUN go mod download

WORKDIR /usr/local/src/app
RUN go mod download

WORKDIR /usr/local/src
COPY shared/ ./shared/
COPY app/ ./app/

WORKDIR /usr/local/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /usr/local/bin/app ./cmd/api/main.go

FROM alpine:3.24

WORKDIR /usr/local
COPY --from=build /usr/local/bin/app /usr/local/bin/app

EXPOSE 8080

CMD [ "/usr/local/bin/app" ]
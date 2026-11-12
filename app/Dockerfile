FROM golang:1.27.3-alpine3.24 AS build

WORKDIR /usr/local/src
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /usr/local/bin/app ./cmd/api/main.go

FROM alpine:3.24

WORKDIR /usr/local
COPY --from=build /usr/local/bin/app /usr/local/bin/app

EXPOSE 8080

CMD [ "/usr/local/bin/app" ]



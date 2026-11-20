module github.com/idbeholdv18/up-code/notifications

go 1.25.1

require (
	github.com/idbeholdv18/up-code/shared v0.0.0
	github.com/rabbitmq/rabbitmq-amqp-go-client v1.5.0
)

replace github.com/idbeholdv18/up-code/shared => ../shared

require (
	github.com/Azure/go-amqp v1.7.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
)

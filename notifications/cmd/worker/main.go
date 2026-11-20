package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/idbeholdv18/up-code/notifications/internal/channel"
	"github.com/idbeholdv18/up-code/notifications/internal/consumer"
	"github.com/idbeholdv18/up-code/notifications/internal/dispatcher"
	"github.com/idbeholdv18/up-code/shared/logger"
	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

const (
	eventsExchangeName = "lms.events"
	notificationQueue  = "notifications.queue"
	routingKeyPattern  = "user.*"
)

func main() {
	logger.Init("notifications-service")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rabbitmqUser := os.Getenv("RABBITMQ_DEFAULT_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_DEFAULT_PASS")
	if rabbitmqUser == "" || rabbitmqPassword == "" {
		log.Fatalf("rabbitmq user and password must be provided by env")
	}

	smtpAddr := os.Getenv("SMTP_ADDR")
	if smtpAddr == "" {
		log.Panicf("smtp addr must be provided by env")

	}
	brokerURI := fmt.Sprintf("amqp://%s:%s@rabbitmq/", rabbitmqUser, rabbitmqPassword)

	env := rmq.NewEnvironment(brokerURI, nil)
	conn, err := env.NewConnection(ctx)
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v", err)
	}
	defer func() {
		_ = env.CloseConnections(context.Background())
	}()

	mgmt := conn.Management()

	_, err = mgmt.DeclareExchange(ctx, &rmq.TopicExchangeSpecification{
		Name: eventsExchangeName,
	})
	if err != nil {
		log.Fatalf("failed to declare exchange %s: %v", eventsExchangeName, err)
	}

	_, err = mgmt.DeclareQueue(ctx, &rmq.QuorumQueueSpecification{
		Name: notificationQueue,
	})
	if err != nil {
		log.Fatalf("failed to declare a queue %s: %v", notificationQueue, err)
	}

	_, err = mgmt.Bind(ctx, &rmq.ExchangeToQueueBindingSpecification{
		SourceExchange:   eventsExchangeName,
		DestinationQueue: notificationQueue,
		BindingKey:       routingKeyPattern,
	})
	if err != nil {
		log.Fatalf("failed to bind queue %s to exchange %s: %v", notificationQueue, eventsExchangeName, err)
	}

	c, err := conn.NewConsumer(ctx, notificationQueue, nil)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}
	defer func() {
		_ = c.Close(context.Background())
	}()

	emailSender := channel.NewSMTPSender(smtpAddr, "noreply@lms.local")
	disp := dispatcher.NewDispatcher(emailSender)

	rmqConsumer := consumer.NewRMQConsumer(c, disp)
	rmqConsumer.Run(ctx)
}

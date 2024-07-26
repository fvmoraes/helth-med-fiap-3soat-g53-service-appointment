package messaging

import (
	"encoding/json"
	"helthmed-appointment/internal/appointment/domain"
	"helthmed-appointment/internal/appointment/infrastructure/rabbitmq"
	"log"

	"github.com/streadway/amqp"
)

func PublishAppointmentCreated(appointment *domain.Appointment) {
	body, err := json.Marshal(appointment)
	if err != nil {
		log.Fatalf("failed to marshal appointment: %v", err)
	}

	err = rabbitmq.Channel.Publish(
		"",
		"appointment_created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}
}

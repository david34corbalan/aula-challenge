package users

import (
	users "uala/pkg/users/ports"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Handler struct {
	UsersService  users.UsersService
	KafkaProducer *kafka.Producer
	KafkaTopic    string
}

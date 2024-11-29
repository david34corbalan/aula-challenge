package handlers

import (
	"uala/cmd/api/database"

	usersHandler "uala/cmd/api/handlers/users"
	usersRepo "uala/pkg/users/repository/postgress"
	usersServices "uala/pkg/users/services"

	tweetsHandler "uala/cmd/api/handlers/tweets"
	tweetsRepo "uala/pkg/tweets/repository/postgress"
	tweetsServices "uala/pkg/tweets/services"

	followHandler "uala/cmd/api/handlers/follow"
	followRepo "uala/pkg/follow/repository/postgress"
	followServices "uala/pkg/follow/services"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type InstanceHandlers struct {
	Users   *usersHandler.Handler
	Tweets  *tweetsHandler.Handler
	Follows *followHandler.Handler
}

func Instance(db database.DataBaseIntance, kafkaProducer *kafka.Producer, kafkaTopic string) *InstanceHandlers {

	// users
	usersRepo := &usersRepo.Repository{
		Client: db.Writer,
	}

	userService := &usersServices.Service{
		Repo: usersRepo,
	}

	usersHandler := &usersHandler.Handler{
		UsersService: userService,
		KafkaProducer: kafkaProducer,
		KafkaTopic:    kafkaTopic,
	}

	// follow
	followsRepo := &followRepo.Repository{
		Client: db.Writer,
	}

	followsServices := &followServices.Service{
		Repo:     followsRepo,
		RepoUser: usersRepo,
	}

	followsHandler := &followHandler.Handler{
		FollowService: followsServices,
	}

	//tweets
	tweetRepo := &tweetsRepo.Repository{
		Client: db.Writer,
	}

	tweetService := &tweetsServices.Service{
		Repo:       tweetRepo,
		RepoUser:   usersRepo,
		RepoFollow: followsRepo,
	}

	tweetsHandler := &tweetsHandler.Handler{
		TweetsService: tweetService,
	}

	return &InstanceHandlers{
		Users:   usersHandler,
		Tweets:  tweetsHandler,
		Follows: followsHandler,
	}
}

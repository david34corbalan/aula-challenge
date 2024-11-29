package routes

import (
	"uala/cmd/api/database"
	"uala/cmd/api/handlers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, db database.DataBaseIntance, kafkaProducer *kafka.Producer, kafkaTopic string) (r *gin.Engine) {
	handlers := handlers.Instance(db, kafkaProducer, kafkaTopic)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")

		users.GET("", handlers.Users.IndexUsers)
		users.GET("/:id", handlers.Users.GetUser)
		users.POST("", handlers.Users.CreateUser)
		users.PUT("/:id", handlers.Users.UpdateUser)
		users.DELETE("/:id", handlers.Users.DeleteUser)

		tweets := v1.Group("/tweets")

		tweets.GET("", handlers.Tweets.IndexTweets)
		tweets.GET("/timeline/:id", handlers.Tweets.TimelieTweet)
		tweets.GET("/:id", handlers.Tweets.GetTweet)
		tweets.POST("", handlers.Tweets.CreateTweet)
		tweets.PUT("/:id", handlers.Tweets.UpdateTweet)
		tweets.DELETE("/:id", handlers.Tweets.DeleteTweet)

		follow := v1.Group("/follow")

		follow.POST("", handlers.Follows.CreateFollow)

	}

	r = router
	return
}

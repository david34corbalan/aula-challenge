package main

import (
	"io"
	"os"
	"uala/cmd/api/database"
	"uala/cmd/api/routes"
	"uala/cmd/api/services"
	"uala/pkg/common/validators"
	follow "uala/pkg/follow/domain"
	tweets "uala/pkg/tweets/domain"
	users "uala/pkg/users/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var (
	Producer *kafka.Producer
	Consumer *kafka.Consumer
)

func main() {
	gin.ForceConsoleColor()
	// no se pueden ver los contenedores
	// Producer, errK := services.NewProducer(services.Config)
	// if errK != nil {
	// 	panic(errK)
	// }
	// defer services.CloseProducer(Producer)
	// services.HandleProducerEvents(Producer)
	// Producer.Flush(15 * 1000) // Tiempo máximo de espera: 15 segundos

	// Consumer, errC := services.NewConsumer(services.Config)
	// if errC != nil {
	// 	panic(errC)
	// }
	// defer services.CloseConsumer(Consumer)
	// go services.ConsumeMessages(Consumer, services.Config.Topic)

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db := database.NewDataBaseIntance()
	db.Connect()
	db.Migrations(users.Users{}, tweets.Tweets{}, follow.Follow{})
	// db.InsertManyRows()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("limit_offset", validators.LimitAndOffset)
	}

	r := routes.InitRoutes(
		gin.Default(func(e *gin.Engine) {
			e.Use(gin.Logger())
			e.Use(gin.Recovery())
			e.Use(gin.ErrorLogger())
			e.Use(cors.New(cors.Config{
				AllowAllOrigins: true,
				AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
			}))
		}),
		*db,
		Producer,
		services.Config.Topic,
	)

	r.Run(":8080")
}

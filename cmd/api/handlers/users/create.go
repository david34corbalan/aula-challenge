package users

import (
	"uala/cmd/api/core"
	users "uala/pkg/users/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateUser(c *gin.Context) {

	var userBinding users.UserCreate
	if err := c.ShouldBindJSON(&userBinding); err != nil {
		response := core.RespondErrorBinding(err, userBinding)
		c.JSON(response.Code, response)
		return
	}

	user, err := h.UsersService.Create(userBinding)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	// Produce a message to Kafka
	// message := []byte(fmt.Sprintf("User %s created", user.Name))
	// err = services.ProduceMessage(h.KafkaProducer, h.KafkaTopic, message)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to produce message"})
	// 	return
	// }

	c.JSON(201, user)
}

package tweets

import (
	"uala/cmd/api/core"
	tweets "uala/pkg/tweets/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateTweet(c *gin.Context) {

	var tweetBinding tweets.TweetsCreate
	if err := c.ShouldBindJSON(&tweetBinding); err != nil {
		response := core.RespondErrorBinding(err, tweetBinding)
		c.JSON(response.Code, response)
		return
	}

	user, err := h.TweetsService.Create(tweetBinding)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(201, user)
}

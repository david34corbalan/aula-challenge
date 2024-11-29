package tweets

import (
	"net/http"
	"strconv"
	"uala/cmd/api/core"
	tweets "uala/pkg/tweets/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) UpdateTweet(c *gin.Context) {
	tweetIdParam := c.Param("id")
	tweetID, err := strconv.Atoi(tweetIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	var tweetBinding tweets.TweetsUpdate
	if err := c.ShouldBindJSON(&tweetBinding); err != nil {
		c.JSON(http.StatusUnprocessableEntity, core.RespondErrorBinding(err, tweetBinding))
		return
	}

	tweet, err := h.TweetsService.Update(tweetID, tweetBinding)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, tweet)
}

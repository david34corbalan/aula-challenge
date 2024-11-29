package tweets

import (
	"strconv"
	"uala/cmd/api/core"

	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteTweet(c *gin.Context) {
	tweetIdParam := c.Param("id")
	tweetID, err := strconv.Atoi(tweetIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	user, err := h.TweetsService.Delete(tweetID)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

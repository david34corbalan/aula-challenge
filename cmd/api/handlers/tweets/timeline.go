package tweets

import (
	"strconv"
	"uala/cmd/api/core"
	"uala/pkg/common"

	"github.com/gin-gonic/gin"
)

func (h Handler) TimelieTweet(c *gin.Context) {
	var params common.QuerysParamsPaginate
	if err := c.ShouldBindQuery(&params); err != nil {
		response := core.RespondErrorBinding(err, params)
		c.JSON(response.Code, response)
		return
	}

	userIdParam := c.Param("id")
	userID, err := strconv.Atoi(userIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	user, err := h.TweetsService.Timeline(userID, params.Limit, params.Offset)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

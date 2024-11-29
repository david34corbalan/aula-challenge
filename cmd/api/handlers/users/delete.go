package users

import (
	"strconv"
	"uala/cmd/api/core"

	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteUser(c *gin.Context) {
	tweetIdParam := c.Param("id")
	tweetID, err := strconv.Atoi(tweetIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	user, err := h.UsersService.Delete(tweetID)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

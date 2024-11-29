package follow

import (
	"uala/cmd/api/core"
	follow "uala/pkg/follow/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateFollow(c *gin.Context) {

	var followBinding follow.FollowUser
	if err := c.ShouldBindJSON(&followBinding); err != nil {
		response := core.RespondErrorBinding(err, followBinding)
		c.JSON(response.Code, response)
		return
	}

	user, err := h.FollowService.Create(followBinding)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(201, user)
}

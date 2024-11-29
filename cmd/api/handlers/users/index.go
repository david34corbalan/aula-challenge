package users

import (
	"uala/cmd/api/core"
	"uala/pkg/common"

	"github.com/gin-gonic/gin"
)

func (h Handler) IndexUsers(c *gin.Context) {

	var params common.QuerysParamsPaginate
	if err := c.ShouldBindQuery(&params); err != nil {
		response := core.RespondErrorBinding(err, params)
		c.JSON(response.Code, response)
		return
	}

	Users, err := h.UsersService.Index(params)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, Users)
}

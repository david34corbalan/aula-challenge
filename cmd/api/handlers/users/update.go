package users

import (
	"net/http"
	"strconv"
	"uala/cmd/api/core"
	users "uala/pkg/users/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) UpdateUser(c *gin.Context) {
	userIdParam := c.Param("id")
	userID, err := strconv.Atoi(userIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	var userBinding users.UserUpdate
	if err := c.ShouldBindJSON(&userBinding); err != nil {
		c.JSON(http.StatusUnprocessableEntity, core.RespondErrorBinding(err, userBinding))
		return
	}

	user, err := h.UsersService.Update(userID, userBinding)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

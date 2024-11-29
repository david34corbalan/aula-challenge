package users

import (
	"fmt"
	"strconv"
	"uala/cmd/api/core"

	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteUser(c *gin.Context) {
	usertIdParam := c.Param("id")
	userID, err := strconv.Atoi(usertIdParam)
	fmt.Println(userID)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	user, err := h.UsersService.Delete(userID)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

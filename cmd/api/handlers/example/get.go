package example

import (
	"github.com/gin-gonic/gin"
)

// GetLeague godoc
// @Summary Get a example
// @Description get a example with id param\
// @Tags example
// @Accept  json
// @Produce  json
// @Param ExampleIdParam body core.ExampleIdParam true "Create example"
// @Success 200 {object} "example: domain.Example"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /example [get]
func (h Handler) GetExample(c *gin.Context) {
	ExampleIdParam := c.Param("id")
	example, err := h.ExampleService.Get(ExampleIdParam)
	if err != nil {
		// core.RespondError(c, err)
		return
	}

	c.JSON(202, example)
}

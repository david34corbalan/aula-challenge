package core_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/core"
	"uala/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type StructError struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

func Test_Reponse_Error(t *testing.T) {
	t.Run("should return code 404", func(t *testing.T) {
		err := common.AppError{Code: 404, Msg: common.ErrCodeInternalServerError}

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		core.RespondError(c, err)

		assert.Equal(t, 404, c.Writer.Status())
		assert.Equal(t, "{\"code\":404,\"message\":\"internal server error\"}", w.Body.String())
		assert.Equal(t, "application/json;", w.Header().Get("Content-Type"))
	})

	t.Run("should return default code 500", func(t *testing.T) {
		err := errors.New("error")

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		core.RespondError(c, err)

		assert.Equal(t, 500, c.Writer.Status())
		assert.Equal(t, "{\"code\":500,\"message\":\"internal server error\"}", w.Body.String())
		assert.Equal(t, "application/json;", w.Header().Get("Content-Type"))
	})
}

func Test_Response_Error_Binding(t *testing.T) {

	t.Run("should return 422 with json error", func(t *testing.T) {

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := `
				{
					"name": "Ezequiel CSUCA",
				}`

		c.Request = httptest.NewRequest("POST", "/test", bytes.NewBuffer([]byte(reqBody)))

		var structBinding StructError
		var res core.ValidationErrorResponse
		var erro error

		if erro = c.ShouldBindJSON(&structBinding); erro != nil {
			res = core.RespondErrorBinding(erro, structBinding)
		}

		assert.Equal(t, 500, res.Code)
		assert.Contains(t, res.Message, "JSON syntax error at byte offset")

	})

	t.Run("should return ValidationErrorResponse", func(t *testing.T) {

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := `
				{
					"name": "Ezequiel CSUCA"
				}`

		c.Request = httptest.NewRequest("POST", "/test", bytes.NewBuffer([]byte(reqBody)))

		var structBinding StructError
		var res core.ValidationErrorResponse
		var erro error

		if erro = c.ShouldBindJSON(&structBinding); erro != nil {
			res = core.RespondErrorBinding(erro, structBinding)
		}

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Message, "Validation Errors")

	})

	t.Run("should return ValidationErrorResponse default ", func(t *testing.T) {

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := `
				{
					"name": "Ezequiel CSUCA"
				}`

		c.Request = httptest.NewRequest("POST", "/test", bytes.NewBuffer([]byte(reqBody)))

		var structBinding StructError
		var res core.ValidationErrorResponse
		var erro error
		if erro = c.ShouldBindJSON(&reqBody); erro != nil {
			res = core.RespondErrorBinding(errors.New("some other error"), &structBinding)
		}

		fmt.Println(res.Message)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Contains(t, res.Message, "JSON syntax error at byte offset some other error")

	})
}

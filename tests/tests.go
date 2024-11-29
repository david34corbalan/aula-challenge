package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Define a type for the setup function
type SetupFunc func(mocks ...any)

// Define a type for the assertion function
type AssertionFuncHandlers func(t *testing.T, w *httptest.ResponseRecorder)

// Define a type for the handler function
type HandlerFunc func(c *gin.Context)

// Define a struct for the test case
type TestCaseHandlers struct {
	Name          string
	Method        string
	Url           string
	Reqbody       string
	Setup         func(...interface{})
	Handler       gin.HandlerFunc
	Assertionfunc func(*testing.T, *httptest.ResponseRecorder)
	Params        gin.Params // Añadimos este campo para los parámetros de la URL
}

func RunTest(t *testing.T, tc TestCaseHandlers, mocks ...interface{}) {
	t.Run(tc.Name, func(subTest *testing.T) {
		if tc.Setup != nil {
			tc.Setup(mocks...)
		}

		req, err := http.NewRequest(tc.Method, tc.Url, bytes.NewBuffer([]byte(tc.Reqbody)))
		if err != nil {
			subTest.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Set the URL parameters if any
		if tc.Params != nil {
			c.Params = tc.Params
		}

		tc.Handler(c)

		if tc.Assertionfunc != nil {
			tc.Assertionfunc(subTest, w)
		}
	})
}

// Define una función de configuración (setup)
type SetupFuncService func(mocks ...interface{})

// Define una función de aserción (assertion)
type AssertionFuncService func(t *testing.T, result interface{}, err error)

// Define una estructura de caso de prueba
type TestCaseService struct {
	Name          string
	Setup         SetupFuncService
	AssertionFunc AssertionFuncService
	TestFunc      func(mocks ...interface{}) (interface{}, error)
}

// Función auxiliar para ejecutar un caso de prueba
func RunTestService(t *testing.T, tc TestCaseService, mocks ...interface{}) {
	t.Run(tc.Name, func(subTest *testing.T) {
		if tc.Setup != nil {
			tc.Setup(mocks...)
		}

		result, err := tc.TestFunc(mocks...)

		if tc.AssertionFunc != nil {
			tc.AssertionFunc(subTest, result, err)
		}
	})
}

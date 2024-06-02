package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

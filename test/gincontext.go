package test

import (
	"bytes"
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

func GetTestGinContextWithParameters(parameters map[string]string) *gin.Context {
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	for key, value := range parameters {
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   key,
			Value: value,
		})
	}

	return ctx
}

func GetTestGinContextWithParametersAndBody(parameters map[string]string, body []byte) *gin.Context {
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	for key, value := range parameters {
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   key,
			Value: value,
		})
	}

	if body != nil {
		request := httptest.NewRequest("POST", "https://test-url.com", bytes.NewReader(body))
		ctx.Request = request
	}

	return ctx
}

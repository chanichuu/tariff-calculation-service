package httphandler

import (
	"bytes"
	"encoding/json"
	"os"
	"tariff-calculation-service/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleGetHealth(t *testing.T) {
	serviceHandler := NewHttpHandler()

	testCtx := test.GetTestGinContext()

	blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: testCtx.Writer}
	testCtx.Writer = blw
	serviceHandler.HandleGetHealth(testCtx)
	statusCode := testCtx.Writer.Status()
	var responseBody string
	err := json.Unmarshal(blw.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, statusCode)
	assert.NotNil(t, responseBody)
	assert.Equal(t, ServiceHealth, responseBody)
}

func Test_HandleGetVersion(t *testing.T) {
	serviceHandler := NewHttpHandler()

	testCtx := test.GetTestGinContext()

	blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: testCtx.Writer}
	testCtx.Writer = blw
	serviceHandler.HandleGetVersion(testCtx)
	statusCode := testCtx.Writer.Status()
	var responseBody string
	err := json.Unmarshal(blw.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, statusCode)
	assert.NotNil(t, responseBody)
	assert.Equal(t, os.Getenv("VERSION"), responseBody)
}

func Test_HandleGetRestVersion(t *testing.T) {
	serviceHandler := NewHttpHandler()

	testCtx := test.GetTestGinContext()

	blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: testCtx.Writer}
	testCtx.Writer = blw
	serviceHandler.HandleGetRestVersion(testCtx)
	statusCode := testCtx.Writer.Status()
	var responseBody string
	err := json.Unmarshal(blw.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, statusCode)
	assert.NotNil(t, responseBody)
	assert.Equal(t, os.Getenv("REST_API_VERSION"), responseBody)
}

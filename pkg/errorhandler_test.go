package pkg

import (
	"bytes"
	"encoding/json"
	"tariff-calculation-service/internal/models"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"testing"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name               string
	ctx                *gin.Context
	expectedStatusCode int
	expectedError      models.Error
	inputError         error
}

func Test_HandleResourceNotFoundAndInternalServerError(t *testing.T) {
	// arrange
	testcases := []testcase{
		{
			"Positive Test ResourceNotFound Error",
			test.GetTestGinContext(),
			404,
			models.NewResourceNotFoundError(),
			errors.New(constants.ResourceNotFound),
		},
		{
			"Positive Test Internal Server Error",
			test.GetTestGinContext(),
			500,
			models.NewInternalServerError(),
			errors.New(constants.InternalServerError),
		},
	}
	// act
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			HandleResourceNotFoundAndInternalServerError(tc.ctx, tc.inputError)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			var actualError models.Error
			err := json.Unmarshal(blw.Body.Bytes(), &actualError)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tc.expectedError, actualError)
		})
	}
}

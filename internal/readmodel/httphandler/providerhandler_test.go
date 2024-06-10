package httphandler

import (
	"bytes"
	"encoding/json"
	"tariff-calculation-service/internal/models"
	repotesting "tariff-calculation-service/internal/readmodel/httphandler/testing"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"testing"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testCaseProviderHandler struct {
	name                 string
	ctx                  *gin.Context
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandleGetProviders(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerGetter := repotesting.NewMockProviderGetter(mockController)

	providerHandler := ProviderHandler{
		ProviderRepo: providerGetter,
	}

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			200,
			&data.Providers,
			func() {
				providerGetter.EXPECT().GetProviders(gomock.Any()).Return(&data.Providers, nil)
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			500,
			models.NewInternalServerError(),
			func() {
				providerGetter.EXPECT().GetProviders(gomock.Any()).Return(&[]models.Provider{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			providerHandler.HandleGetProviders(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualProviders *[]models.Provider
				err := json.Unmarshal(blw.Body.Bytes(), &actualProviders)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualProviders)
				assert.GreaterOrEqual(t, 1, len(*actualProviders))
				assert.Equal(t, tc.expectedResponse, actualProviders)
			} else {
				var actualError models.Error
				err := json.Unmarshal(blw.Body.Bytes(), &actualError)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualError)
				assert.Equal(t, tc.expectedResponse, actualError)
			}
		})
	}
}

func Test_HandleGetProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerGetter := repotesting.NewMockProviderGetter(mockController)

	providerHandler := ProviderHandler{
		ProviderRepo: providerGetter,
	}

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			200,
			&data.Provider,
			func() {
				providerGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&data.Provider, nil)
			},
		},
		{
			"Negative Test Not Found",
			test.GetTestGinContext(),
			404,
			models.NewResourceNotFoundError(),
			func() {
				providerGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&models.Provider{}, errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			500,
			models.NewInternalServerError(),
			func() {
				providerGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&models.Provider{}, errors.New(constants.InternalServerError))
			},
		},
	}

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			providerHandler.HandleGetProvider(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualProvider *models.Provider
				err := json.Unmarshal(blw.Body.Bytes(), &actualProvider)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualProvider)
				assert.Equal(t, tc.expectedResponse, actualProvider)
			} else {
				var actualError models.Error
				err := json.Unmarshal(blw.Body.Bytes(), &actualError)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualError)
				assert.Equal(t, tc.expectedResponse, actualError)
			}
		})
	}
}

package writehandlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	repotesting "tariff-calculation-service/internal/writemodel/writehandlers/testing"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"tariff-calculation-service/test/mocks"
	"tariff-calculation-service/tools"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

type depsProvider struct {
	repo      ProviderWriter
	validator interfaces.Validator
}

type testCasePWH struct {
	name                 string
	ctx                  *gin.Context
	deps                 depsProvider
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandlePostProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerRepo := repotesting.NewMockProviderWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCasePWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: mockValidator},
			201,
			&data.Provider,
			func() { providerRepo.EXPECT().CreateProvider(gomock.Any(), gomock.Any()).Return(&data.Provider, nil) },
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				providerRepo.EXPECT().CreateProvider(gomock.Any(), gomock.Any()).Return(nil, errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerWriteHandler := ProviderHandler{ProviderWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			providerWriteHandler.HandlePostProvider(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 201 {
				actualProvider := models.Provider{}
				err := json.Unmarshal(blw.Body.Bytes(), &actualProvider)
				if err != nil {
					t.Fail()
				}

				assert.Equal(t, tc.expectedResponse, actualProvider)
			} else {
				var actualError models.Error
				err := json.Unmarshal(blw.Body.Bytes(), &actualError)
				if err != nil {
					t.Fail()
				}
				assert.Equal(t, tc.expectedResponse, actualError)
			}
		})
	}
}

func Test_HandlePutProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerRepo := repotesting.NewMockProviderWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCasePWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: mockValidator},
			204,
			nil,
			func() { providerRepo.EXPECT().UpdateProvider(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				providerRepo.EXPECT().UpdateProvider(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				providerRepo.EXPECT().UpdateProvider(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerWriteHandler := ProviderHandler{ProviderWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			providerWriteHandler.HandlePutProvider(tc.ctx)
			statusCode := tc.ctx.Writer.Status()
			if statusCode != 204 {
				var actualError models.Error
				err := json.Unmarshal(blw.Body.Bytes(), &actualError)
				if err != nil {
					t.Fail()
				}
				assert.Equal(t, tc.expectedResponse, actualError)
			}

			assert.Equal(t, tc.expectedResponseCode, statusCode)
		})
	}
}

func Test_HandleDeleteProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerRepo := repotesting.NewMockProviderWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCasePWH{
		{
			"Positive Test",
			test.GetTestGinContext(),
			depsProvider{repo: providerRepo, validator: mockValidator},
			204,
			nil,
			func() { providerRepo.EXPECT().DeleteProvider(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContext(),
			depsProvider{repo: providerRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				providerRepo.EXPECT().DeleteProvider(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			depsProvider{repo: providerRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				providerRepo.EXPECT().DeleteProvider(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerWriteHandler := ProviderHandler{ProviderWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			providerWriteHandler.HandleDeleteProvider(tc.ctx)
			statusCode := tc.ctx.Writer.Status()
			if statusCode != 204 {
				var actualError models.Error
				err := json.Unmarshal(blw.Body.Bytes(), &actualError)
				if err != nil {
					t.Fail()
				}
				assert.Equal(t, tc.expectedResponse, actualError)
			}

			assert.Equal(t, tc.expectedResponseCode, statusCode)
		})
	}
}

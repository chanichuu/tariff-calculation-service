package httphandler

import (
	"bytes"
	"encoding/json"
	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	repotesting "tariff-calculation-service/internal/readmodel/httphandler/testing"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"tariff-calculation-service/test/mocks"
	"testing"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type dependenciesProviderHandler struct {
	repo      ProviderGetter
	validator interfaces.Validator
}

type testCaseProviderHandler struct {
	name                 string
	ctx                  *gin.Context
	deps                 dependenciesProviderHandler
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandleGetProviders(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockProviderGetter := repotesting.NewMockProviderGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			200,
			&data.Providers,
			func() {
				mockProviderGetter.EXPECT().GetProviders(gomock.Any()).Return(&data.Providers, nil)
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockProviderGetter.EXPECT().GetProviders(gomock.Any()).Return(&[]models.Provider{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerHandler := ProviderHandler{
				ProviderRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
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

func Test_HandleGetProviders_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockProviderGetter := repotesting.NewMockProviderGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			200,
			&data.Providers,
			func() {
				mockProviderGetter.EXPECT().GetProviders(gomock.Any()).Return(&data.Providers, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerHandler := ProviderHandler{
				ProviderRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
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

	mockProviderGetter := repotesting.NewMockProviderGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			200,
			&data.Provider,
			func() {
				mockProviderGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&data.Provider, nil)
			},
		},
		{
			"Negative Test Not Found",
			test.GetTestGinContext(),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				mockProviderGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&models.Provider{}, errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockProviderGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&models.Provider{}, errors.New(constants.InternalServerError))
			},
		},
	}

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerHandler := ProviderHandler{
				ProviderRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
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

func Test_HandleGetProvider_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockProviderGetter := repotesting.NewMockProviderGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testCases := []testCaseProviderHandler{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidator},
			200,
			&data.Provider,
			func() {
				mockProviderGetter.EXPECT().GetProvider(gomock.Any(), gomock.Any()).Return(&data.Provider, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestProviderId}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestIdInvalid}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test PartitionId and Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestIdInvalid}),
			dependenciesProviderHandler{repo: mockProviderGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerHandler := ProviderHandler{
				ProviderRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
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

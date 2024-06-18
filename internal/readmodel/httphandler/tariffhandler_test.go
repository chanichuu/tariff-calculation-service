package httphandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	repotesting "tariff-calculation-service/internal/readmodel/httphandler/testing"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"tariff-calculation-service/test/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type dependenciesTariffHandler struct {
	repo      TariffGetter
	validator interfaces.Validator
}

type testCaseTariffHandler struct {
	name                 string
	ctx                  *gin.Context
	deps                 dependenciesTariffHandler
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_GetTariffs(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTariffGetter := repotesting.NewMockTariffGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseTariffHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			200,
			&data.Tariffs,
			func() {
				mockTariffGetter.EXPECT().GetTariffs(gomock.Any()).Return(&data.Tariffs, nil)
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockTariffGetter.EXPECT().GetTariffs(gomock.Any()).Return(&[]models.Tariff{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffHandler := TariffHandler{
				TariffRepo: tc.deps.repo,
				Validator:  tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			tariffHandler.HandleGetTariffs(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualTariffs *[]models.Tariff
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariffs)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualTariffs)
				assert.GreaterOrEqual(t, 1, len(*actualTariffs))
				assert.Equal(t, tc.expectedResponse, actualTariffs)
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

func Test_GetTariffs_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTariffGetter := repotesting.NewMockTariffGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testCases := []testCaseTariffHandler{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			200,
			&data.Tariffs,
			func() {
				mockTariffGetter.EXPECT().GetTariffs(gomock.Any()).Return(&data.Tariffs, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffHandler := TariffHandler{
				TariffRepo: tc.deps.repo,
				Validator:  tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			tariffHandler.HandleGetTariffs(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualTariffs *[]models.Tariff
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariffs)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualTariffs)
				assert.GreaterOrEqual(t, 1, len(*actualTariffs))
				assert.Equal(t, tc.expectedResponse, actualTariffs)
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

func Test_GetTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTariffGetter := repotesting.NewMockTariffGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseTariffHandler{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			200,
			&data.Tariff,
			func() {
				mockTariffGetter.EXPECT().GetTariff(gomock.Any(), gomock.Any()).Return(&data.Tariff, nil)
			},
		},
		{
			"Negative Test Not Found",
			test.GetTestGinContext(),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				mockTariffGetter.EXPECT().GetTariff(gomock.Any(), gomock.Any()).Return(&models.Tariff{}, errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockTariffGetter.EXPECT().GetTariff(gomock.Any(), gomock.Any()).Return(&models.Tariff{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffHandler := TariffHandler{
				TariffRepo: tc.deps.repo,
				Validator:  tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			tariffHandler.HandleGetTariff(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualTariff *models.Tariff
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariff)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualTariff)
				assert.Equal(t, tc.expectedResponse, actualTariff)
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

func Test_GetTariff_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTariffGetter := repotesting.NewMockTariffGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testCases := []testCaseTariffHandler{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidator},
			200,
			&data.Tariff,
			func() {
				mockTariffGetter.EXPECT().GetTariff(gomock.Any(), gomock.Any()).Return(&data.Tariff, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestTariffId}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestIdInvalid}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test PartitionId and Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestIdInvalid}),
			dependenciesTariffHandler{repo: mockTariffGetter, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}
	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffHandler := TariffHandler{
				TariffRepo: tc.deps.repo,
				Validator:  tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			tariffHandler.HandleGetTariff(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualTariff *models.Tariff
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariff)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualTariff)
				assert.Equal(t, tc.expectedResponse, actualTariff)
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

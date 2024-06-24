package writehandlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
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

type depsTariff struct {
	repo      TariffWriter
	validator interfaces.Validator
}

type testCaseTWH struct {
	name                 string
	ctx                  *gin.Context
	deps                 depsTariff
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandlePostTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	tariffRepo := repotesting.NewMockTariffWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseTWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			201,
			&data.Tariff,
			func() { tariffRepo.EXPECT().CreateTariff(gomock.Any(), gomock.Any()).Return(&data.Tariff, nil) },
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				tariffRepo.EXPECT().CreateTariff(gomock.Any(), gomock.Any()).Return(nil, errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffWriteHandler := TariffHandler{TariffWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			tariffWriteHandler.HandlePostTariff(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 201 {
				actualTariff := models.Tariff{}
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariff)
				if err != nil {
					t.Fail()
				}

				assert.Equal(t, tc.expectedResponse, actualTariff)
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

func Test_HandlePostTariff_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	tariffRepo := repotesting.NewMockTariffWriter(mockController)
	validator := mocks.NewValidatorPathPositive(mockController)
	validatorNegative := mocks.NewValidatorPathNegative(mockController)

	tariffEmptyName := data.Tariff
	tariffEmptyName.Name = ""

	tariffNameLenExceeded := data.Tariff
	tariffNameLenExceeded.Name = strings.Repeat("a", 65)

	tariffEmptyCurrency := data.Tariff
	tariffEmptyCurrency.Currency = ""

	tariffCurrencyInvalid := data.Tariff
	tariffCurrencyInvalid.Currency = "Invalid-Currency"

	tariffEmptyValidFrom := data.Tariff
	tariffEmptyValidFrom.ValidFrom = ""

	tariffValidFromInvalid := data.Tariff
	tariffValidFromInvalid.ValidFrom = "01/01/2023"

	tariffEmptyValidTo := data.Tariff
	tariffEmptyValidTo.ValidTo = ""

	tariffValidToInvalid := data.Tariff
	tariffValidToInvalid.ValidTo = "01/01/2023"

	tariffInvalidFixedPrice := data.Tariff
	tariffInvalidFixedPrice.FixedTariff = models.FixedTariff{PricePerUnit: -1}

	tariffInvalidStartTimeHourly := data.TariffInvalidHourlyStartTime
	tariffInvalidStartTimeHourly.DynamicTariff.HourlyTariffs[0].StartTime = "01/01/2023"

	tariffInvalidValidDaysHourly := data.TariffInvalidHourlyValidDays
	tariffInvalidValidDaysHourly.DynamicTariff.HourlyTariffs[0].ValidDays = []uint8{1, 2, 3, 4, 5, 6, 7, 7}

	testCases := []testCaseTWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: validator},
			201,
			&data.Tariff,
			func() { tariffRepo.EXPECT().CreateTariff(gomock.Any(), gomock.Any()).Return(&data.Tariff, nil) },
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: validatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty Name",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffEmptyName))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Name Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffNameLenExceeded))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty Currency",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffEmptyCurrency))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Currency", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid Currency",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffCurrencyInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Currency", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty ValidFrom",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffEmptyValidFrom))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidFrom", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid ValidFrom",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffValidFromInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidFrom", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty ValidTo",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffEmptyValidTo))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidTo", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid ValidTo",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffValidToInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidTo", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid FixedPrice",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffInvalidFixedPrice))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PricePerUnit", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty StartTime Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyStartTime))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"StartTime", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid StartTime Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffInvalidStartTimeHourly))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"StartTime", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff ValidDays Max Value Exceeded Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyValidDays))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidDays", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff ValidDays Max Len Exceeded Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(tariffInvalidValidDaysHourly))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidDays", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid PricePerUnit Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyPricePerUnit))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PricePerUnit", ""}})),
			func() {
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffWriteHandler := TariffHandler{TariffWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			tariffWriteHandler.HandlePostTariff(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 201 {
				actualTariff := models.Tariff{}
				err := json.Unmarshal(blw.Body.Bytes(), &actualTariff)
				if err != nil {
					t.Fail()
				}

				assert.Equal(t, tc.expectedResponse, actualTariff)
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

func Test_HandlePutTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	tariffRepo := repotesting.NewMockTariffWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseTWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			204,
			nil,
			func() { tariffRepo.EXPECT().UpdateTariff(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				tariffRepo.EXPECT().UpdateTariff(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				tariffRepo.EXPECT().UpdateTariff(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffWriteHandler := TariffHandler{TariffWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			tariffWriteHandler.HandlePutTariff(tc.ctx)
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

func Test_HandlePutTariff_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	tariffRepo := repotesting.NewMockTariffWriter(mockController)
	validator := mocks.NewValidatorPathPositive(mockController)
	validatorNegative := mocks.NewValidatorPathNegative(mockController)

	tariffEmptyName := data.Tariff
	tariffEmptyName.Name = ""

	tariffNameLenExceeded := data.Tariff
	tariffNameLenExceeded.Name = strings.Repeat("a", 65)

	tariffEmptyCurrency := data.Tariff
	tariffEmptyCurrency.Currency = ""

	tariffCurrencyInvalid := data.Tariff
	tariffCurrencyInvalid.Currency = "Invalid-Currency"

	tariffEmptyValidFrom := data.Tariff
	tariffEmptyValidFrom.ValidFrom = ""

	tariffValidFromInvalid := data.Tariff
	tariffValidFromInvalid.ValidFrom = "01/01/2023"

	tariffEmptyValidTo := data.Tariff
	tariffEmptyValidTo.ValidTo = ""

	tariffValidToInvalid := data.Tariff
	tariffValidToInvalid.ValidTo = "01/01/2023"

	tariffInvalidFixedPrice := data.Tariff
	tariffInvalidFixedPrice.FixedTariff = models.FixedTariff{PricePerUnit: -1}

	tariffInvalidStartTimeHourly := data.TariffInvalidHourlyStartTime
	tariffInvalidStartTimeHourly.DynamicTariff.HourlyTariffs[0].StartTime = "01/01/2023"

	tariffInvalidValidDaysHourly := data.TariffInvalidHourlyValidDays
	tariffInvalidValidDaysHourly.DynamicTariff.HourlyTariffs[0].ValidDays = []uint8{1, 2, 3, 4, 5, 6, 7, 7}

	testCases := []testCaseTWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: validator},
			204,
			nil,
			func() { tariffRepo.EXPECT().UpdateTariff(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.Tariff))),
			depsTariff{repo: tariffRepo, validator: validatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty Name",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffEmptyName))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Name Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffNameLenExceeded))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty Currency",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffEmptyCurrency))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Currency", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid Currency",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffCurrencyInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Currency", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty ValidFrom",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffEmptyValidFrom))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidFrom", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid ValidFrom",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffValidFromInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidFrom", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty ValidTo",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffEmptyValidTo))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidTo", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid ValidTo",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffValidToInvalid))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidTo", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid FixedPrice",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffInvalidFixedPrice))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PricePerUnit", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Empty StartTime Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyStartTime))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"StartTime", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid StartTime Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffInvalidStartTimeHourly))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"StartTime", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff ValidDays Max Value Exceeded Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyValidDays))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidDays", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff ValidDays Max Len Exceeded Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(tariffInvalidValidDaysHourly))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"ValidDays", ""}})),
			func() {
			},
		},
		{
			"Negative Test Tariff Invalid PricePerUnit Hourly",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestTariffId}, tools.GetFirstValue(json.Marshal(data.TariffInvalidHourlyPricePerUnit))),
			depsTariff{repo: tariffRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PricePerUnit", ""}})),
			func() {
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffWriteHandler := TariffHandler{TariffWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			tariffWriteHandler.HandlePutTariff(tc.ctx)
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

func Test_HandleDeleteTariff(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	tariffRepo := repotesting.NewMockTariffWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseTWH{
		{
			"Positive Test",
			test.GetTestGinContext(),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			204,
			nil,
			func() { tariffRepo.EXPECT().DeleteTariff(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContext(),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				tariffRepo.EXPECT().DeleteTariff(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			depsTariff{repo: tariffRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				tariffRepo.EXPECT().DeleteTariff(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tariffWriteHandler := TariffHandler{TariffWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			tariffWriteHandler.HandleDeleteTariff(tc.ctx)
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

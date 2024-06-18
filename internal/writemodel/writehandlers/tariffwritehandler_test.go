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

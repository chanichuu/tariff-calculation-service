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

type dependencies struct {
	repo      ContractWriter
	validator interfaces.Validator
}

type testCaseCWH struct {
	name                 string
	ctx                  *gin.Context
	deps                 dependencies
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandlePostContract(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	contractRepo := repotesting.NewMockContractWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseCWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Contract))),
			dependencies{repo: contractRepo, validator: mockValidator},
			201,
			&data.Contract,
			func() { contractRepo.EXPECT().CreateContract(gomock.Any(), gomock.Any()).Return(&data.Contract, nil) },
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Contract))),
			dependencies{repo: contractRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				contractRepo.EXPECT().CreateContract(gomock.Any(), gomock.Any()).Return(nil, errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contractWriteHandler := ContractWriteHandler{ContractWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			contractWriteHandler.HandlePostContract(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 201 {
				actualContract := models.Contract{}
				err := json.Unmarshal(blw.Body.Bytes(), &actualContract)
				if err != nil {
					t.Fail()
				}

				assert.Equal(t, tc.expectedResponse, actualContract)
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

func Test_HandlePutContract(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	contractRepo := repotesting.NewMockContractWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseCWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestContractId}, tools.GetFirstValue(json.Marshal(data.Contract))),
			dependencies{repo: contractRepo, validator: mockValidator},
			204,
			nil,
			func() { contractRepo.EXPECT().UpdateContract(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestContractId}, tools.GetFirstValue(json.Marshal(data.Contract))),
			dependencies{repo: contractRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				contractRepo.EXPECT().UpdateContract(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestContractId}, tools.GetFirstValue(json.Marshal(data.Contract))),
			dependencies{repo: contractRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				contractRepo.EXPECT().UpdateContract(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contractWriteHandler := ContractWriteHandler{ContractWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			contractWriteHandler.HandlePutContract(tc.ctx)
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

func Test_HandleDeleteContract(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	contractRepo := repotesting.NewMockContractWriter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testCases := []testCaseCWH{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependencies{repo: contractRepo, validator: mockValidator},
			204,
			nil,
			func() { contractRepo.EXPECT().DeleteContract(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test Resource Not Found",
			test.GetTestGinContext(),
			dependencies{repo: contractRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				contractRepo.EXPECT().DeleteContract(gomock.Any(), gomock.Any()).Return(errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependencies{repo: contractRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				contractRepo.EXPECT().DeleteContract(gomock.Any(), gomock.Any()).Return(errors.New(constants.InternalServerError))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contractWriteHandler := ContractWriteHandler{ContractWriter: tc.deps.repo, Validator: tc.deps.validator}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw

			contractWriteHandler.HandleDeleteContract(tc.ctx)
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

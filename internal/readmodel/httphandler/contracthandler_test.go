package httphandler

import (
	"encoding/json"
	"tariff-calculation-service/internal/interfaces"
	"tariff-calculation-service/internal/models"
	repoMocks "tariff-calculation-service/internal/readmodel/httphandler/testing"
	"tariff-calculation-service/pkg/constants"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"tariff-calculation-service/test/mocks"
	"testing"

	"bytes"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type dependencies struct {
	repo      ContractGetter
	validator interfaces.Validator
}

type testcase struct {
	name                 string
	ctx                  *gin.Context
	deps                 dependencies
	expectedResponseCode int
	expectedResponse     any
	mockFunc             func()
}

func Test_HandleGetContracts(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractRepo := repoMocks.NewMockContractGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testcases := []testcase{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			200,
			&data.Contracts,
			func() {
				mockContractRepo.EXPECT().GetContracts(gomock.Any()).Return(&data.Contracts, nil)
			},
		},
		{
			"Negative Test",
			test.GetTestGinContext(),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockContractRepo.EXPECT().GetContracts(gomock.Any()).Return(&[]models.Contract{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			contractHandler := ContractHandler{
				ContractRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}

			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			contractHandler.HandleGetContracts(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualContracts *[]models.Contract
				err := json.Unmarshal(blw.Body.Bytes(), &actualContracts)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualContracts)
				assert.GreaterOrEqual(t, 1, len(*actualContracts))
				assert.Equal(t, tc.expectedResponse, actualContracts)
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

func Test_HandleGetContracts_Validation(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractRepo := repoMocks.NewMockContractGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testcases := []testcase{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId}),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			200,
			&data.Contracts,
			func() {
				mockContractRepo.EXPECT().GetContracts(gomock.Any()).AnyTimes().Return(&data.Contracts, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid}),
			dependencies{repo: mockContractRepo, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}
	// act
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			contractHandler := ContractHandler{
				ContractRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			contractHandler.HandleGetContracts(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualContracts *[]models.Contract
				err := json.Unmarshal(blw.Body.Bytes(), &actualContracts)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualContracts)
				assert.GreaterOrEqual(t, 1, len(*actualContracts))
				assert.Equal(t, tc.expectedResponse, actualContracts)
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

func Test_HandleGetContract(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractRepo := repoMocks.NewMockContractGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)

	testcases := []testcase{
		{
			"Positive Test",
			test.GetTestGinContext(),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			200,
			&data.Contract,
			func() { mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).Return(&data.Contract, nil) },
		},
		{
			"Negative Test Contract Not Found",
			test.GetTestGinContext(),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			404,
			models.NewResourceNotFoundError(),
			func() {
				mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).Return(&models.Contract{}, errors.New(constants.ResourceNotFound))
			},
		},
		{
			"Negative Test Internal Server Error",
			test.GetTestGinContext(),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			500,
			models.NewInternalServerError(),
			func() {
				mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).Return(&models.Contract{}, errors.New(constants.InternalServerError))
			},
		},
	}
	// act
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			contractHandler := ContractHandler{
				ContractRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			contractHandler.HandleGetContract(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualContract *models.Contract
				err := json.Unmarshal(blw.Body.Bytes(), &actualContract)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualContract)
				assert.Equal(t, tc.expectedResponse, actualContract)
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

func Test_HandleGetContract_Validation(t *testing.T) {
	// arrange
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractRepo := repoMocks.NewMockContractGetter(mockController)
	mockValidator := mocks.NewValidatorPathPositive(mockController)
	mockValidatorNegative := mocks.NewValidatorPathNegative(mockController)

	testcases := []testcase{
		{
			"Positive Test",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestContractId}),
			dependencies{repo: mockContractRepo, validator: mockValidator},
			200,
			&data.Contract,
			func() {
				mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).AnyTimes().Return(&data.Contract, nil)
			},
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestContractId}),
			dependencies{repo: mockContractRepo, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
				mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).AnyTimes().Return(&data.Contract, nil)
			},
		},
		{
			"Negative Test Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestIdInvalid}),
			dependencies{repo: mockContractRepo, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
				mockContractRepo.EXPECT().GetContract(gomock.Any(), gomock.Any()).AnyTimes().Return(&data.Contract, nil)

			},
		},
		{
			"Negative Test PartitionId and Id Invalid",
			test.GetTestGinContextWithParameters(map[string]string{"PartitionId": data.TestIdInvalid, "Id": data.TestIdInvalid}),
			dependencies{repo: mockContractRepo, validator: mockValidatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
	}
	// act
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			contractHandler := ContractHandler{
				ContractRepo: tc.deps.repo,
				Validator:    tc.deps.validator,
			}
			tc.mockFunc()
			blw := &test.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: tc.ctx.Writer}
			tc.ctx.Writer = blw
			contractHandler.HandleGetContract(tc.ctx)
			statusCode := tc.ctx.Writer.Status()

			// assert
			assert.Equal(t, tc.expectedResponseCode, statusCode)
			if statusCode == 200 {
				var actualContract *models.Contract
				err := json.Unmarshal(blw.Body.Bytes(), &actualContract)
				if err != nil {
					t.Fail()
				}
				assert.NotNil(t, actualContract)
				assert.Equal(t, tc.expectedResponse, actualContract)
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

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

func Test_HandlePostProvider_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerRepo := repotesting.NewMockProviderWriter(mockController)
	validator := mocks.NewValidatorPathPositive(mockController)
	validatorNegative := mocks.NewValidatorPathNegative(mockController)

	providerNameEmpty := data.Provider
	providerNameEmpty.Name = ""

	providerNameLenExceeded := data.Provider
	providerNameLenExceeded.Name = strings.Repeat("a", 65)

	providerEmailInvalid := data.Provider
	providerEmailInvalid.Email = "Invalid-Email"

	providerAddressStreetEmpty := data.Provider
	providerAddressStreetEmpty.Address.Street = ""

	providerAddressStreetLenExceeded := data.Provider
	providerAddressStreetLenExceeded.Address.Street = strings.Repeat("a", 65)

	providerAddressPostalCodeEmpty := data.Provider
	providerAddressPostalCodeEmpty.Address.PostalCode = ""

	providerAddressPostalCodeLenExceeded := data.Provider
	providerAddressPostalCodeLenExceeded.Address.PostalCode = strings.Repeat("a", 65)

	providerAddressCityEmpty := data.Provider
	providerAddressCityEmpty.Address.City = ""

	providerAddressCityLenExceeded := data.Provider
	providerAddressCityLenExceeded.Address.City = strings.Repeat("a", 65)

	providerAddressCountryCodeEmpty := data.Provider
	providerAddressCountryCodeEmpty.Address.CountryCode = ""

	providerAddressCountryCodeInvalid := data.Provider
	providerAddressCountryCodeInvalid.Address.CountryCode = "Invalid-CC"

	testCases := []testCasePWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: validator},
			201,
			&data.Provider,
			func() { providerRepo.EXPECT().CreateProvider(gomock.Any(), gomock.Any()).Return(&data.Provider, nil) },
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: validatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Name",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerNameEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Name Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerNameLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider EMail Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerEmailInvalid))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Email", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address Street",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressStreetEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Street", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address Street Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressStreetLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Street", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address PostalCode",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressPostalCodeEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PostalCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address PostalCode Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressPostalCodeLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PostalCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address City",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressCityEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"City", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address City Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressCityLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"City", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address CountryCode",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressCountryCodeEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"CountryCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address CountryCode Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId}, tools.GetFirstValue(json.Marshal(providerAddressCountryCodeInvalid))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"CountryCode", ""}})),
			func() {
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

func Test_HandlePutProvider_Validation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	providerRepo := repotesting.NewMockProviderWriter(mockController)
	validator := mocks.NewValidatorPathPositive(mockController)
	validatorNegative := mocks.NewValidatorPathNegative(mockController)

	providerNameEmpty := data.Provider
	providerNameEmpty.Name = ""

	providerNameLenExceeded := data.Provider
	providerNameLenExceeded.Name = strings.Repeat("a", 65)

	providerEmailInvalid := data.Provider
	providerEmailInvalid.Email = "Invalid-Email"

	providerAddressStreetEmpty := data.Provider
	providerAddressStreetEmpty.Address.Street = ""

	providerAddressStreetLenExceeded := data.Provider
	providerAddressStreetLenExceeded.Address.Street = strings.Repeat("a", 65)

	providerAddressPostalCodeEmpty := data.Provider
	providerAddressPostalCodeEmpty.Address.PostalCode = ""

	providerAddressPostalCodeLenExceeded := data.Provider
	providerAddressPostalCodeLenExceeded.Address.PostalCode = strings.Repeat("a", 65)

	providerAddressCityEmpty := data.Provider
	providerAddressCityEmpty.Address.City = ""

	providerAddressCityLenExceeded := data.Provider
	providerAddressCityLenExceeded.Address.City = strings.Repeat("a", 65)

	providerAddressCountryCodeEmpty := data.Provider
	providerAddressCountryCodeEmpty.Address.CountryCode = ""

	providerAddressCountryCodeInvalid := data.Provider
	providerAddressCountryCodeInvalid.Address.CountryCode = "Invalid-CC"

	testCases := []testCasePWH{
		{
			"Positive Test",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: validator},
			204,
			nil,
			func() { providerRepo.EXPECT().UpdateProvider(gomock.Any(), gomock.Any()).Return(nil) },
		},
		{
			"Negative Test PartitionId Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(data.Provider))),
			depsProvider{repo: providerRepo, validator: validatorNegative},
			400,
			models.NewBadRequestFieldValidationError(errors.New("ValidationError")),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Name",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerNameEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Name Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerNameLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Name", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider EMail Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerEmailInvalid))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Email", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address Street",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressStreetEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Street", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address Street Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressStreetLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"Street", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address PostalCode",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressPostalCodeEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PostalCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address PostalCode Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressPostalCodeLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"PostalCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address City",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressCityEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"City", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address City Max Length Exceeded",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressCityLenExceeded))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"City", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Empty Address CountryCode",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressCountryCodeEmpty))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"CountryCode", ""}})),
			func() {
			},
		},
		{
			"Negative Test Provider Address CountryCode Invalid",
			test.GetTestGinContextWithParametersAndBody(map[string]string{"PartitionId": data.TestPartitionId, "Id": data.TestProviderId}, tools.GetFirstValue(json.Marshal(providerAddressCountryCodeInvalid))),
			depsProvider{repo: providerRepo, validator: validator},
			400,
			models.NewBadRequestFieldValidationError(data.FieldValidationError([][2]string{{"CountryCode", ""}})),
			func() {
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

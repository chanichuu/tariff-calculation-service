package validation

import (
	"fmt"
	"tariff-calculation-service/test"
	"tariff-calculation-service/test/data"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type args[T any] struct {
	ctx    *gin.Context
	object *T
}

type testCaseValidation[T any] struct {
	name               string
	args               args[T]
	wantErr            assert.ErrorAssertionFunc
	expectedStatusCode int
	expectedObject     T
}

func Test_ValidateAndSetPathParams_PartitionId(t *testing.T) {
	validator := NewValidator()

	var testPathPartitionId = PartitionId{}

	testCases := []testCaseValidation[PartitionId]{
		{
			"Positive Test",
			args[PartitionId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestPartitionId}), &testPathPartitionId},
			assert.NoError,
			200,
			PartitionId{
				PartitionId: data.TestPartitionId,
			},
		},
		{
			"Negative Test Invalid PartitionId",
			args[PartitionId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestIdInvalid}), &testPathPartitionId},
			assert.Error,
			400,
			PartitionId{
				PartitionId: data.TestIdInvalid,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.wantErr(t, validator.ValidateAndSetPathParams(tc.args.ctx, tc.args.object), fmt.Sprintf("ValidateAndSetPathParams(%v,%v)", tc.args.ctx, tc.args.object))

			assert.Equal(t, tc.expectedStatusCode, tc.args.ctx.Writer.Status())
			assert.Equal(t, tc.expectedObject, *tc.args.object)
		})
	}
}

func Test_ValidateAndSetPathParams_PartitionIdWithId(t *testing.T) {
	validator := NewValidator()

	var testPathPartitionId = PartitionIdWithId{}

	testCases := []testCaseValidation[PartitionIdWithId]{
		{
			"Positive Test",
			args[PartitionIdWithId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestPartitionId, "id": data.TestTariffId}), &testPathPartitionId},
			assert.NoError,
			200,
			PartitionIdWithId{
				PartitionId: data.TestPartitionId,
				Id:          data.TestTariffId,
			},
		},
		{
			"Negative Test Invalid PartitionId",
			args[PartitionIdWithId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestIdInvalid, "id": data.TestTariffId}), &testPathPartitionId},
			assert.Error,
			400,
			PartitionIdWithId{
				PartitionId: data.TestIdInvalid,
				Id:          data.TestTariffId,
			},
		},
		{
			"Negative Test Invalid Id",
			args[PartitionIdWithId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestPartitionId, "id": data.TestIdInvalid}), &testPathPartitionId},
			assert.Error,
			400,
			PartitionIdWithId{
				PartitionId: data.TestPartitionId,
				Id:          data.TestIdInvalid,
			},
		},
		{
			"Negative Test Invalid PartitionId and Id",
			args[PartitionIdWithId]{test.GetTestGinContextWithParameters(map[string]string{"partitionId": data.TestIdInvalid, "id": data.TestIdInvalid}), &testPathPartitionId},
			assert.Error,
			400,
			PartitionIdWithId{
				PartitionId: data.TestIdInvalid,
				Id:          data.TestIdInvalid,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.wantErr(t, validator.ValidateAndSetPathParams(tc.args.ctx, tc.args.object), fmt.Sprintf("ValidateAndSetPathParams(%v,%v)", tc.args.ctx, tc.args.object))

			assert.Equal(t, tc.expectedStatusCode, tc.args.ctx.Writer.Status())
			assert.Equal(t, tc.expectedObject, *tc.args.object)
		})
	}
}

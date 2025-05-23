// Code generated by MockGen. DO NOT EDIT.
// Source: contractwritehandler.go
//
// Generated by this command:
//
//	mockgen -source=contractwritehandler.go -destination=testing/contractwritehandler_mocks.go -package=testing ContractWriter
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"
	models "tariff-calculation-service/internal/models"

	gomock "go.uber.org/mock/gomock"
)

// MockContractWriter is a mock of ContractWriter interface.
type MockContractWriter struct {
	ctrl     *gomock.Controller
	recorder *MockContractWriterMockRecorder
}

// MockContractWriterMockRecorder is the mock recorder for MockContractWriter.
type MockContractWriterMockRecorder struct {
	mock *MockContractWriter
}

// NewMockContractWriter creates a new mock instance.
func NewMockContractWriter(ctrl *gomock.Controller) *MockContractWriter {
	mock := &MockContractWriter{ctrl: ctrl}
	mock.recorder = &MockContractWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractWriter) EXPECT() *MockContractWriterMockRecorder {
	return m.recorder
}

// CreateContract mocks base method.
func (m *MockContractWriter) CreateContract(partitionId string, contract models.Contract) (*models.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContract", partitionId, contract)
	ret0, _ := ret[0].(*models.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContract indicates an expected call of CreateContract.
func (mr *MockContractWriterMockRecorder) CreateContract(partitionId, contract any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContract", reflect.TypeOf((*MockContractWriter)(nil).CreateContract), partitionId, contract)
}

// DeleteContract mocks base method.
func (m *MockContractWriter) DeleteContract(partitionId, contractId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContract", partitionId, contractId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContract indicates an expected call of DeleteContract.
func (mr *MockContractWriterMockRecorder) DeleteContract(partitionId, contractId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContract", reflect.TypeOf((*MockContractWriter)(nil).DeleteContract), partitionId, contractId)
}

// UpdateContract mocks base method.
func (m *MockContractWriter) UpdateContract(partitionId string, contract models.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContract", partitionId, contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContract indicates an expected call of UpdateContract.
func (mr *MockContractWriterMockRecorder) UpdateContract(partitionId, contract any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContract", reflect.TypeOf((*MockContractWriter)(nil).UpdateContract), partitionId, contract)
}

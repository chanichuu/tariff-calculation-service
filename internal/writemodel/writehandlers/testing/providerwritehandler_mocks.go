// Code generated by MockGen. DO NOT EDIT.
// Source: providerwritehandler.go
//
// Generated by this command:
//
//	mockgen -source=providerwritehandler.go -destination=testing/providerwritehandler_mocks.go -package=testing ProviderWriter
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"
	models "tariff-calculation-service/internal/models"

	gomock "go.uber.org/mock/gomock"
)

// MockProviderWriter is a mock of ProviderWriter interface.
type MockProviderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockProviderWriterMockRecorder
}

// MockProviderWriterMockRecorder is the mock recorder for MockProviderWriter.
type MockProviderWriterMockRecorder struct {
	mock *MockProviderWriter
}

// NewMockProviderWriter creates a new mock instance.
func NewMockProviderWriter(ctrl *gomock.Controller) *MockProviderWriter {
	mock := &MockProviderWriter{ctrl: ctrl}
	mock.recorder = &MockProviderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderWriter) EXPECT() *MockProviderWriterMockRecorder {
	return m.recorder
}

// CreateProvider mocks base method.
func (m *MockProviderWriter) CreateProvider(partitionId string, provider models.Provider) (*models.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProvider", partitionId, provider)
	ret0, _ := ret[0].(*models.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProvider indicates an expected call of CreateProvider.
func (mr *MockProviderWriterMockRecorder) CreateProvider(partitionId, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProvider", reflect.TypeOf((*MockProviderWriter)(nil).CreateProvider), partitionId, provider)
}

// DeleteProvider mocks base method.
func (m *MockProviderWriter) DeleteProvider(partitionId, providerId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProvider", partitionId, providerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProvider indicates an expected call of DeleteProvider.
func (mr *MockProviderWriterMockRecorder) DeleteProvider(partitionId, providerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProvider", reflect.TypeOf((*MockProviderWriter)(nil).DeleteProvider), partitionId, providerId)
}

// UpdateProvider mocks base method.
func (m *MockProviderWriter) UpdateProvider(partitionId string, provider models.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProvider", partitionId, provider)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProvider indicates an expected call of UpdateProvider.
func (mr *MockProviderWriterMockRecorder) UpdateProvider(partitionId, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProvider", reflect.TypeOf((*MockProviderWriter)(nil).UpdateProvider), partitionId, provider)
}

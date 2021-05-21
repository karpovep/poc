// Code generated by MockGen. DO NOT EDIT.
// Source: poc/repository (interfaces: IRepository)

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	nodes "poc/protos/nodes"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// ResetActiveIsoNodeId mocks base method.
func (m *MockIRepository) ResetActiveIsoNodeId(arg0 *nodes.ISO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetActiveIsoNodeId", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetActiveIsoNodeId indicates an expected call of ResetActiveIsoNodeId.
func (mr *MockIRepositoryMockRecorder) ResetActiveIsoNodeId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetActiveIsoNodeId", reflect.TypeOf((*MockIRepository)(nil).ResetActiveIsoNodeId), arg0)
}

// Start mocks base method.
func (m *MockIRepository) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockIRepositoryMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockIRepository)(nil).Start))
}

// Stop mocks base method.
func (m *MockIRepository) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockIRepositoryMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIRepository)(nil).Stop))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: poc/utils (interfaces: IUtils)

// Package utils_mock is a generated GoMock package.
package utils_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUtils is a mock of IUtils interface.
type MockIUtils struct {
	ctrl     *gomock.Controller
	recorder *MockIUtilsMockRecorder
}

// MockIUtilsMockRecorder is the mock recorder for MockIUtils.
type MockIUtilsMockRecorder struct {
	mock *MockIUtils
}

// NewMockIUtils creates a new mock instance.
func NewMockIUtils(ctrl *gomock.Controller) *MockIUtils {
	mock := &MockIUtils{ctrl: ctrl}
	mock.recorder = &MockIUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUtils) EXPECT() *MockIUtilsMockRecorder {
	return m.recorder
}

// GenerateTimeUuid mocks base method.
func (m *MockIUtils) GenerateTimeUuid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTimeUuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateTimeUuid indicates an expected call of GenerateTimeUuid.
func (mr *MockIUtilsMockRecorder) GenerateTimeUuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTimeUuid", reflect.TypeOf((*MockIUtils)(nil).GenerateTimeUuid))
}

// GenerateUuid mocks base method.
func (m *MockIUtils) GenerateUuid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateUuid indicates an expected call of GenerateUuid.
func (mr *MockIUtilsMockRecorder) GenerateUuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUuid", reflect.TypeOf((*MockIUtils)(nil).GenerateUuid))
}

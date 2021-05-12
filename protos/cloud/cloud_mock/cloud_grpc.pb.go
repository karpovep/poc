// Code generated by MockGen. DO NOT EDIT.
// Source: poc/protos/cloud (interfaces: Cloud_SubscribeServer)

// Package cloud_mock is a generated GoMock package.
package cloud_mock

import (
	context "context"
	cloud "poc/protos/cloud"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockCloud_SubscribeServer is a mock of Cloud_SubscribeServer interface.
type MockCloud_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockCloud_SubscribeServerMockRecorder
}

// MockCloud_SubscribeServerMockRecorder is the mock recorder for MockCloud_SubscribeServer.
type MockCloud_SubscribeServerMockRecorder struct {
	mock *MockCloud_SubscribeServer
}

// NewMockCloud_SubscribeServer creates a new mock instance.
func NewMockCloud_SubscribeServer(ctrl *gomock.Controller) *MockCloud_SubscribeServer {
	mock := &MockCloud_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockCloud_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud_SubscribeServer) EXPECT() *MockCloud_SubscribeServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockCloud_SubscribeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockCloud_SubscribeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockCloud_SubscribeServer) Recv() (*cloud.CloudObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*cloud.CloudObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockCloud_SubscribeServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockCloud_SubscribeServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockCloud_SubscribeServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockCloud_SubscribeServer) Send(arg0 *cloud.CloudObject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockCloud_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockCloud_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockCloud_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockCloud_SubscribeServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCloud_SubscribeServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockCloud_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockCloud_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockCloud_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockCloud_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockCloud_SubscribeServer)(nil).SetTrailer), arg0)
}
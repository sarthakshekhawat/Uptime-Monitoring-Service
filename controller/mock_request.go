// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sarthakshekhawat/Uptime-Monitoring-Service/controller (interfaces: RequestInterface)

// Package controller is a generated GoMock package.
package controller

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRequestInterface is a mock of RequestInterface interface.
type MockRequestInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRequestInterfaceMockRecorder
}

// MockRequestInterfaceMockRecorder is the mock recorder for MockRequestInterface.
type MockRequestInterfaceMockRecorder struct {
	mock *MockRequestInterface
}

// NewMockRequestInterface creates a new mock instance.
func NewMockRequestInterface(ctrl *gomock.Controller) *MockRequestInterface {
	mock := &MockRequestInterface{ctrl: ctrl}
	mock.recorder = &MockRequestInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestInterface) EXPECT() *MockRequestInterfaceMockRecorder {
	return m.recorder
}

// httpRequest mocks base method.
func (m *MockRequestInterface) httpRequest(arg0 DataBase) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "httpRequest", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// httpRequest indicates an expected call of httpRequest.
func (mr *MockRequestInterfaceMockRecorder) httpRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "httpRequest", reflect.TypeOf((*MockRequestInterface)(nil).httpRequest), arg0)
}

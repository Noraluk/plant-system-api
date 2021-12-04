// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/controllers/pump/interface/pump.go

// Package mock_pumpctrlitf is a generated GoMock package.
package mock_pumpctrlitf

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockPumpController is a mock of PumpController interface.
type MockPumpController struct {
	ctrl     *gomock.Controller
	recorder *MockPumpControllerMockRecorder
}

// MockPumpControllerMockRecorder is the mock recorder for MockPumpController.
type MockPumpControllerMockRecorder struct {
	mock *MockPumpController
}

// NewMockPumpController creates a new mock instance.
func NewMockPumpController(ctrl *gomock.Controller) *MockPumpController {
	mock := &MockPumpController{ctrl: ctrl}
	mock.recorder = &MockPumpControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPumpController) EXPECT() *MockPumpControllerMockRecorder {
	return m.recorder
}

// ActivePump mocks base method.
func (m *MockPumpController) ActivePump(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivePump", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivePump indicates an expected call of ActivePump.
func (mr *MockPumpControllerMockRecorder) ActivePump(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivePump", reflect.TypeOf((*MockPumpController)(nil).ActivePump), c)
}

// AskPump mocks base method.
func (m *MockPumpController) AskPump(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskPump", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AskPump indicates an expected call of AskPump.
func (mr *MockPumpControllerMockRecorder) AskPump(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskPump", reflect.TypeOf((*MockPumpController)(nil).AskPump), c)
}

// GetPump mocks base method.
func (m *MockPumpController) GetPump(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPump", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetPump indicates an expected call of GetPump.
func (mr *MockPumpControllerMockRecorder) GetPump(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPump", reflect.TypeOf((*MockPumpController)(nil).GetPump), c)
}

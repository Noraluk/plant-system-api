// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/routes/pump/interface/pump.go

// Package mock_pumprtitf is a generated GoMock package.
package mock_pumprtitf

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPumpRoute is a mock of PumpRoute interface.
type MockPumpRoute struct {
	ctrl     *gomock.Controller
	recorder *MockPumpRouteMockRecorder
}

// MockPumpRouteMockRecorder is the mock recorder for MockPumpRoute.
type MockPumpRouteMockRecorder struct {
	mock *MockPumpRoute
}

// NewMockPumpRoute creates a new mock instance.
func NewMockPumpRoute(ctrl *gomock.Controller) *MockPumpRoute {
	mock := &MockPumpRoute{ctrl: ctrl}
	mock.recorder = &MockPumpRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPumpRoute) EXPECT() *MockPumpRouteMockRecorder {
	return m.recorder
}

// SetRoutes mocks base method.
func (m *MockPumpRoute) SetRoutes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRoutes")
}

// SetRoutes indicates an expected call of SetRoutes.
func (mr *MockPumpRouteMockRecorder) SetRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoutes", reflect.TypeOf((*MockPumpRoute)(nil).SetRoutes))
}

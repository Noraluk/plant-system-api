// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/routes/health/interface/health.go

// Package mock_healthrtitf is a generated GoMock package.
package mock_healthrtitf

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHealthRoute is a mock of HealthRoute interface.
type MockHealthRoute struct {
	ctrl     *gomock.Controller
	recorder *MockHealthRouteMockRecorder
}

// MockHealthRouteMockRecorder is the mock recorder for MockHealthRoute.
type MockHealthRouteMockRecorder struct {
	mock *MockHealthRoute
}

// NewMockHealthRoute creates a new mock instance.
func NewMockHealthRoute(ctrl *gomock.Controller) *MockHealthRoute {
	mock := &MockHealthRoute{ctrl: ctrl}
	mock.recorder = &MockHealthRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthRoute) EXPECT() *MockHealthRouteMockRecorder {
	return m.recorder
}

// SetRoutes mocks base method.
func (m *MockHealthRoute) SetRoutes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRoutes")
}

// SetRoutes indicates an expected call of SetRoutes.
func (mr *MockHealthRouteMockRecorder) SetRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoutes", reflect.TypeOf((*MockHealthRoute)(nil).SetRoutes))
}

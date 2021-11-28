// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/services/pump/interface/pump.go

// Package mock_pumpserviceitf is a generated GoMock package.
package mock_pumpserviceitf

import (
	pumpmodel "plant-system-api/api/models/pump"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPumpService is a mock of PumpService interface.
type MockPumpService struct {
	ctrl     *gomock.Controller
	recorder *MockPumpServiceMockRecorder
}

// MockPumpServiceMockRecorder is the mock recorder for MockPumpService.
type MockPumpServiceMockRecorder struct {
	mock *MockPumpService
}

// NewMockPumpService creates a new mock instance.
func NewMockPumpService(ctrl *gomock.Controller) *MockPumpService {
	mock := &MockPumpService{ctrl: ctrl}
	mock.recorder = &MockPumpServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPumpService) EXPECT() *MockPumpServiceMockRecorder {
	return m.recorder
}

// ActivePump mocks base method.
func (m *MockPumpService) ActivePump(pump *pumpmodel.Pump) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivePump", pump)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivePump indicates an expected call of ActivePump.
func (mr *MockPumpServiceMockRecorder) ActivePump(pump interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivePump", reflect.TypeOf((*MockPumpService)(nil).ActivePump), pump)
}

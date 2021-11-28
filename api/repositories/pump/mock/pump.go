// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/repositories/pump/interface/pump.go

// Package mock_pumprepoitf is a generated GoMock package.
package mock_pumprepoitf

import (
	pumpmodel "plant-system-api/api/models/pump"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPumpRepository is a mock of PumpRepository interface.
type MockPumpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPumpRepositoryMockRecorder
}

// MockPumpRepositoryMockRecorder is the mock recorder for MockPumpRepository.
type MockPumpRepositoryMockRecorder struct {
	mock *MockPumpRepository
}

// NewMockPumpRepository creates a new mock instance.
func NewMockPumpRepository(ctrl *gomock.Controller) *MockPumpRepository {
	mock := &MockPumpRepository{ctrl: ctrl}
	mock.recorder = &MockPumpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPumpRepository) EXPECT() *MockPumpRepositoryMockRecorder {
	return m.recorder
}

// ActivePump mocks base method.
func (m *MockPumpRepository) ActivePump(pump *pumpmodel.Pump) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivePump", pump)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivePump indicates an expected call of ActivePump.
func (mr *MockPumpRepositoryMockRecorder) ActivePump(pump interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivePump", reflect.TypeOf((*MockPumpRepository)(nil).ActivePump), pump)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_taxiFare is a generated GoMock package.
package mock_taxiFare

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	taxiFare "taxiFare/internal/app/repository/taxiFare"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetDetailFare mocks base method
func (m *MockRepository) GetDetailFare(param taxiFare.Param) (taxiFare.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailFare", param)
	ret0, _ := ret[0].(taxiFare.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailFare indicates an expected call of GetDetailFare
func (mr *MockRepositoryMockRecorder) GetDetailFare(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailFare", reflect.TypeOf((*MockRepository)(nil).GetDetailFare), param)
}

// GetData mocks base method
func (m *MockRepository) GetData(param taxiFare.Param) ([]taxiFare.ResponseRedis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", param)
	ret0, _ := ret[0].([]taxiFare.ResponseRedis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData
func (mr *MockRepositoryMockRecorder) GetData(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockRepository)(nil).GetData), param)
}

// StoreData mocks base method
func (m *MockRepository) StoreData(param taxiFare.Param) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreData", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreData indicates an expected call of StoreData
func (mr *MockRepositoryMockRecorder) StoreData(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreData", reflect.TypeOf((*MockRepository)(nil).StoreData), param)
}

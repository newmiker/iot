// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/micky/Documents/playground/go/iot/internal/aggregation/storage.go

// Package mock_aggregation is a generated GoMock package.
package mock_aggregation

import (
	aggregation "github.com/Michaellqa/iot/internal/aggregation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockStorage) Write(record aggregation.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockStorageMockRecorder) Write(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorage)(nil).Write), record)
}
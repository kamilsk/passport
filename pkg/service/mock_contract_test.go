// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/passport/pkg/service (interfaces: Storage)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "github.com/kamilsk/passport/pkg/domain"
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

// StoreFingerprint mocks base method
func (m *MockStorage) StoreFingerprint(arg0 context.Context, arg1 domain.Fingerprint) (domain.Fingerprint, error) {
	ret := m.ctrl.Call(m, "StoreFingerprint", arg0, arg1)
	ret0, _ := ret[0].(domain.Fingerprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreFingerprint indicates an expected call of StoreFingerprint
func (mr *MockStorageMockRecorder) StoreFingerprint(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreFingerprint", reflect.TypeOf((*MockStorage)(nil).StoreFingerprint), arg0, arg1)
}

// UUID mocks base method
func (m *MockStorage) UUID(arg0 context.Context) (domain.UUID, error) {
	ret := m.ctrl.Call(m, "UUID", arg0)
	ret0, _ := ret[0].(domain.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UUID indicates an expected call of UUID
func (mr *MockStorageMockRecorder) UUID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockStorage)(nil).UUID), arg0)
}

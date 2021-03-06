// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockTag is a mock of Tag interface
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockTag) GetAll() ([]*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTagMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTag)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockTag) GetByID(id int) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTagMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTag)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockTag) Create(name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTagMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTag)(nil).Create), name)
}

// UpdateByID mocks base method
func (m *MockTag) UpdateByID(id int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockTagMockRecorder) UpdateByID(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockTag)(nil).UpdateByID), id, name)
}

// DeleteByID mocks base method
func (m *MockTag) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockTagMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTag)(nil).DeleteByID), id)
}

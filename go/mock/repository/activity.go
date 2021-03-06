// Code generated by MockGen. DO NOT EDIT.
// Source: activity.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockActivityRepository is a mock of ActivityRepository interface
type MockActivityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockActivityRepositoryMockRecorder
}

// MockActivityRepositoryMockRecorder is the mock recorder for MockActivityRepository
type MockActivityRepositoryMockRecorder struct {
	mock *MockActivityRepository
}

// NewMockActivityRepository creates a new mock instance
func NewMockActivityRepository(ctrl *gomock.Controller) *MockActivityRepository {
	mock := &MockActivityRepository{ctrl: ctrl}
	mock.recorder = &MockActivityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActivityRepository) EXPECT() *MockActivityRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockActivityRepository) FindAll() ([]*entity.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockActivityRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockActivityRepository)(nil).FindAll))
}

// FindByID mocks base method
func (m *MockActivityRepository) FindByID(id int) (*entity.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*entity.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockActivityRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockActivityRepository)(nil).FindByID), id)
}

// FindUpcoming mocks base method
func (m *MockActivityRepository) FindUpcoming() ([]*entity.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUpcoming")
	ret0, _ := ret[0].([]*entity.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUpcoming indicates an expected call of FindUpcoming
func (mr *MockActivityRepositoryMockRecorder) FindUpcoming() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUpcoming", reflect.TypeOf((*MockActivityRepository)(nil).FindUpcoming))
}

// FindByNotify mocks base method
func (m *MockActivityRepository) FindByNotify() ([]*entity.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNotify")
	ret0, _ := ret[0].([]*entity.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNotify indicates an expected call of FindByNotify
func (mr *MockActivityRepositoryMockRecorder) FindByNotify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNotify", reflect.TypeOf((*MockActivityRepository)(nil).FindByNotify))
}

// Create mocks base method
func (m *MockActivityRepository) Create(arg0 *entity.Activity) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockActivityRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActivityRepository)(nil).Create), arg0)
}

// UpdateByID mocks base method
func (m *MockActivityRepository) UpdateByID(arg0 *entity.Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockActivityRepositoryMockRecorder) UpdateByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockActivityRepository)(nil).UpdateByID), arg0)
}

// DeleteByID mocks base method
func (m *MockActivityRepository) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockActivityRepositoryMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockActivityRepository)(nil).DeleteByID), id)
}

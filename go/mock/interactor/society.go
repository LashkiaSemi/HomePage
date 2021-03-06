// Code generated by MockGen. DO NOT EDIT.
// Source: society.go

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockSocietyInteractor is a mock of SocietyInteractor interface
type MockSocietyInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockSocietyInteractorMockRecorder
}

// MockSocietyInteractorMockRecorder is the mock recorder for MockSocietyInteractor
type MockSocietyInteractorMockRecorder struct {
	mock *MockSocietyInteractor
}

// NewMockSocietyInteractor creates a new mock instance
func NewMockSocietyInteractor(ctrl *gomock.Controller) *MockSocietyInteractor {
	mock := &MockSocietyInteractor{ctrl: ctrl}
	mock.recorder = &MockSocietyInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSocietyInteractor) EXPECT() *MockSocietyInteractorMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockSocietyInteractor) GetAll() ([]*entity.Society, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Society)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockSocietyInteractorMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSocietyInteractor)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockSocietyInteractor) GetByID(id int) (*entity.Society, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Society)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockSocietyInteractorMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSocietyInteractor)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockSocietyInteractor) Create(title, author, society, award, date string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, author, society, award, date)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSocietyInteractorMockRecorder) Create(title, author, society, award, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSocietyInteractor)(nil).Create), title, author, society, award, date)
}

// UpdateByID mocks base method
func (m *MockSocietyInteractor) UpdateByID(id int, title, author, society, award, date string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, title, author, society, award, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockSocietyInteractorMockRecorder) UpdateByID(id, title, author, society, award, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockSocietyInteractor)(nil).UpdateByID), id, title, author, society, award, date)
}

// DeleteByID mocks base method
func (m *MockSocietyInteractor) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockSocietyInteractorMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockSocietyInteractor)(nil).DeleteByID), id)
}

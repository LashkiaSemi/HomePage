// Code generated by MockGen. DO NOT EDIT.
// Source: lecture.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockLecture is a mock of Lecture interface
type MockLecture struct {
	ctrl     *gomock.Controller
	recorder *MockLectureMockRecorder
}

// MockLectureMockRecorder is the mock recorder for MockLecture
type MockLectureMockRecorder struct {
	mock *MockLecture
}

// NewMockLecture creates a new mock instance
func NewMockLecture(ctrl *gomock.Controller) *MockLecture {
	mock := &MockLecture{ctrl: ctrl}
	mock.recorder = &MockLectureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLecture) EXPECT() *MockLectureMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockLecture) GetAll() ([]*entity.Lecture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Lecture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockLectureMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockLecture)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockLecture) GetByID(id int) (*entity.Lecture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Lecture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockLectureMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLecture)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockLecture) Create(title, file, comment string, activation int, author *entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, file, comment, activation, author)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockLectureMockRecorder) Create(title, file, comment, activation, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLecture)(nil).Create), title, file, comment, activation, author)
}

// UpdateByID mocks base method
func (m *MockLecture) UpdateByID(id int, title, file, comment string, activation int, author *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, title, file, comment, activation, author)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockLectureMockRecorder) UpdateByID(id, title, file, comment, activation, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockLecture)(nil).UpdateByID), id, title, file, comment, activation, author)
}

// DeleteByID mocks base method
func (m *MockLecture) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockLectureMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockLecture)(nil).DeleteByID), id)
}

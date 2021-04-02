// Code generated by MockGen. DO NOT EDIT.
// Source: lecture.go

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockLectureInteractor is a mock of LectureInteractor interface
type MockLectureInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockLectureInteractorMockRecorder
}

// MockLectureInteractorMockRecorder is the mock recorder for MockLectureInteractor
type MockLectureInteractorMockRecorder struct {
	mock *MockLectureInteractor
}

// NewMockLectureInteractor creates a new mock instance
func NewMockLectureInteractor(ctrl *gomock.Controller) *MockLectureInteractor {
	mock := &MockLectureInteractor{ctrl: ctrl}
	mock.recorder = &MockLectureInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLectureInteractor) EXPECT() *MockLectureInteractorMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockLectureInteractor) GetAll() ([]*entity.Lecture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Lecture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockLectureInteractorMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockLectureInteractor)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockLectureInteractor) GetByID(id int) (*entity.Lecture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Lecture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockLectureInteractorMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLectureInteractor)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockLectureInteractor) Create(studentID, title, file, comment string, activation int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", studentID, title, file, comment, activation)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockLectureInteractorMockRecorder) Create(studentID, title, file, comment, activation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLectureInteractor)(nil).Create), studentID, title, file, comment, activation)
}

// UpdateByID mocks base method
func (m *MockLectureInteractor) UpdateByID(id int, studentID, title, file, comment string, activation int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, studentID, title, file, comment, activation)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockLectureInteractorMockRecorder) UpdateByID(id, studentID, title, file, comment, activation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockLectureInteractor)(nil).UpdateByID), id, studentID, title, file, comment, activation)
}

// DeleteByID mocks base method
func (m *MockLectureInteractor) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockLectureInteractorMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockLectureInteractor)(nil).DeleteByID), id)
}

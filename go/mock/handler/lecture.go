// Code generated by MockGen. DO NOT EDIT.
// Source: lecture.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockLectureHandler is a mock of LectureHandler interface
type MockLectureHandler struct {
	ctrl     *gomock.Controller
	recorder *MockLectureHandlerMockRecorder
}

// MockLectureHandlerMockRecorder is the mock recorder for MockLectureHandler
type MockLectureHandlerMockRecorder struct {
	mock *MockLectureHandler
}

// NewMockLectureHandler creates a new mock instance
func NewMockLectureHandler(ctrl *gomock.Controller) *MockLectureHandler {
	mock := &MockLectureHandler{ctrl: ctrl}
	mock.recorder = &MockLectureHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLectureHandler) EXPECT() *MockLectureHandlerMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockLectureHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAll", w, r)
}

// GetAll indicates an expected call of GetAll
func (mr *MockLectureHandlerMockRecorder) GetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockLectureHandler)(nil).GetAll), w, r)
}

// Create mocks base method
func (m *MockLectureHandler) Create(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", w, r)
}

// Create indicates an expected call of Create
func (mr *MockLectureHandlerMockRecorder) Create(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLectureHandler)(nil).Create), w, r)
}

// UpdateByID mocks base method
func (m *MockLectureHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateByID", w, r)
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockLectureHandlerMockRecorder) UpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockLectureHandler)(nil).UpdateByID), w, r)
}

// DeleteByID mocks base method
func (m *MockLectureHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteByID", w, r)
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockLectureHandlerMockRecorder) DeleteByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockLectureHandler)(nil).DeleteByID), w, r)
}

// AdminGetAll mocks base method
func (m *MockLectureHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetAll", w, r)
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockLectureHandlerMockRecorder) AdminGetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockLectureHandler)(nil).AdminGetAll), w, r)
}

// AdminGetByID mocks base method
func (m *MockLectureHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetByID", w, r)
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockLectureHandlerMockRecorder) AdminGetByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockLectureHandler)(nil).AdminGetByID), w, r)
}

// AdminCreate mocks base method
func (m *MockLectureHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminCreate", w, r)
}

// AdminCreate indicates an expected call of AdminCreate
func (mr *MockLectureHandlerMockRecorder) AdminCreate(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminCreate", reflect.TypeOf((*MockLectureHandler)(nil).AdminCreate), w, r)
}

// AdminUpdateByID mocks base method
func (m *MockLectureHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminUpdateByID", w, r)
}

// AdminUpdateByID indicates an expected call of AdminUpdateByID
func (mr *MockLectureHandlerMockRecorder) AdminUpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminUpdateByID", reflect.TypeOf((*MockLectureHandler)(nil).AdminUpdateByID), w, r)
}

// AdminDeleteByID mocks base method
func (m *MockLectureHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminDeleteByID", w, r)
}

// AdminDeleteByID indicates an expected call of AdminDeleteByID
func (mr *MockLectureHandlerMockRecorder) AdminDeleteByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminDeleteByID", reflect.TypeOf((*MockLectureHandler)(nil).AdminDeleteByID), w, r)
}

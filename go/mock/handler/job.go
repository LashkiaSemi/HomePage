// Code generated by MockGen. DO NOT EDIT.
// Source: job.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockJobHandler is a mock of JobHandler interface
type MockJobHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJobHandlerMockRecorder
}

// MockJobHandlerMockRecorder is the mock recorder for MockJobHandler
type MockJobHandlerMockRecorder struct {
	mock *MockJobHandler
}

// NewMockJobHandler creates a new mock instance
func NewMockJobHandler(ctrl *gomock.Controller) *MockJobHandler {
	mock := &MockJobHandler{ctrl: ctrl}
	mock.recorder = &MockJobHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobHandler) EXPECT() *MockJobHandlerMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockJobHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAll", w, r)
}

// GetAll indicates an expected call of GetAll
func (mr *MockJobHandlerMockRecorder) GetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockJobHandler)(nil).GetAll), w, r)
}

// AdminCreate mocks base method
func (m *MockJobHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminCreate", w, r)
}

// AdminCreate indicates an expected call of AdminCreate
func (mr *MockJobHandlerMockRecorder) AdminCreate(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminCreate", reflect.TypeOf((*MockJobHandler)(nil).AdminCreate), w, r)
}

// AdminUpdateByID mocks base method
func (m *MockJobHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminUpdateByID", w, r)
}

// AdminUpdateByID indicates an expected call of AdminUpdateByID
func (mr *MockJobHandlerMockRecorder) AdminUpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminUpdateByID", reflect.TypeOf((*MockJobHandler)(nil).AdminUpdateByID), w, r)
}

// AdminGetAll mocks base method
func (m *MockJobHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetAll", w, r)
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockJobHandlerMockRecorder) AdminGetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockJobHandler)(nil).AdminGetAll), w, r)
}

// AdminGetByID mocks base method
func (m *MockJobHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetByID", w, r)
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockJobHandlerMockRecorder) AdminGetByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockJobHandler)(nil).AdminGetByID), w, r)
}

// AdminDeleteByID mocks base method
func (m *MockJobHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminDeleteByID", w, r)
}

// AdminDeleteByID indicates an expected call of AdminDeleteByID
func (mr *MockJobHandlerMockRecorder) AdminDeleteByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminDeleteByID", reflect.TypeOf((*MockJobHandler)(nil).AdminDeleteByID), w, r)
}
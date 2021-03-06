// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockUserHandler is a mock of UserHandler interface
type MockUserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlerMockRecorder
}

// MockUserHandlerMockRecorder is the mock recorder for MockUserHandler
type MockUserHandlerMockRecorder struct {
	mock *MockUserHandler
}

// NewMockUserHandler creates a new mock instance
func NewMockUserHandler(ctrl *gomock.Controller) *MockUserHandler {
	mock := &MockUserHandler{ctrl: ctrl}
	mock.recorder = &MockUserHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserHandler) EXPECT() *MockUserHandlerMockRecorder {
	return m.recorder
}

// GetAllGroupByGrade mocks base method
func (m *MockUserHandler) GetAllGroupByGrade(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllGroupByGrade", w, r)
}

// GetAllGroupByGrade indicates an expected call of GetAllGroupByGrade
func (mr *MockUserHandlerMockRecorder) GetAllGroupByGrade(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGroupByGrade", reflect.TypeOf((*MockUserHandler)(nil).GetAllGroupByGrade), w, r)
}

// GetByID mocks base method
func (m *MockUserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetByID", w, r)
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserHandlerMockRecorder) GetByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserHandler)(nil).GetByID), w, r)
}

// UpdateByID mocks base method
func (m *MockUserHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateByID", w, r)
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockUserHandlerMockRecorder) UpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockUserHandler)(nil).UpdateByID), w, r)
}

// UpdatePasswordByStudentID mocks base method
func (m *MockUserHandler) UpdatePasswordByStudentID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePasswordByStudentID", w, r)
}

// UpdatePasswordByStudentID indicates an expected call of UpdatePasswordByStudentID
func (mr *MockUserHandlerMockRecorder) UpdatePasswordByStudentID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePasswordByStudentID", reflect.TypeOf((*MockUserHandler)(nil).UpdatePasswordByStudentID), w, r)
}

// Login mocks base method
func (m *MockUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", w, r)
}

// Login indicates an expected call of Login
func (mr *MockUserHandlerMockRecorder) Login(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserHandler)(nil).Login), w, r)
}

// Logout mocks base method
func (m *MockUserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Logout", w, r)
}

// Logout indicates an expected call of Logout
func (mr *MockUserHandlerMockRecorder) Logout(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserHandler)(nil).Logout), w, r)
}

// AdminLogin mocks base method
func (m *MockUserHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminLogin", w, r)
}

// AdminLogin indicates an expected call of AdminLogin
func (mr *MockUserHandlerMockRecorder) AdminLogin(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminLogin", reflect.TypeOf((*MockUserHandler)(nil).AdminLogin), w, r)
}

// AdminGetAll mocks base method
func (m *MockUserHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetAll", w, r)
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockUserHandlerMockRecorder) AdminGetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockUserHandler)(nil).AdminGetAll), w, r)
}

// AdminGetByID mocks base method
func (m *MockUserHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetByID", w, r)
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockUserHandlerMockRecorder) AdminGetByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockUserHandler)(nil).AdminGetByID), w, r)
}

// AdminCreate mocks base method
func (m *MockUserHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminCreate", w, r)
}

// AdminCreate indicates an expected call of AdminCreate
func (mr *MockUserHandlerMockRecorder) AdminCreate(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminCreate", reflect.TypeOf((*MockUserHandler)(nil).AdminCreate), w, r)
}

// AdminUpdateByID mocks base method
func (m *MockUserHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminUpdateByID", w, r)
}

// AdminUpdateByID indicates an expected call of AdminUpdateByID
func (mr *MockUserHandlerMockRecorder) AdminUpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminUpdateByID", reflect.TypeOf((*MockUserHandler)(nil).AdminUpdateByID), w, r)
}

// AdminDeleteByID mocks base method
func (m *MockUserHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminDeleteByID", w, r)
}

// AdminDeleteByID indicates an expected call of AdminDeleteByID
func (mr *MockUserHandlerMockRecorder) AdminDeleteByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminDeleteByID", reflect.TypeOf((*MockUserHandler)(nil).AdminDeleteByID), w, r)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: society.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	gomock "github.com/golang/mock/gomock"
	controller "homepage/pkg/interface/controller"
	reflect "reflect"
)

// MockSocietyController is a mock of SocietyController interface
type MockSocietyController struct {
	ctrl     *gomock.Controller
	recorder *MockSocietyControllerMockRecorder
}

// MockSocietyControllerMockRecorder is the mock recorder for MockSocietyController
type MockSocietyControllerMockRecorder struct {
	mock *MockSocietyController
}

// NewMockSocietyController creates a new mock instance
func NewMockSocietyController(ctrl *gomock.Controller) *MockSocietyController {
	mock := &MockSocietyController{ctrl: ctrl}
	mock.recorder = &MockSocietyControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSocietyController) EXPECT() *MockSocietyControllerMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockSocietyController) GetAll() (*controller.SocietiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*controller.SocietiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockSocietyControllerMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSocietyController)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockSocietyController) GetByID(id int) (*controller.SocietyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*controller.SocietyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockSocietyControllerMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSocietyController)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockSocietyController) Create(title, author, society, award, date string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, author, society, award, date)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSocietyControllerMockRecorder) Create(title, author, society, award, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSocietyController)(nil).Create), title, author, society, award, date)
}

// UpdateByID mocks base method
func (m *MockSocietyController) UpdateByID(id int, title, author, society, award, date string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, title, author, society, award, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockSocietyControllerMockRecorder) UpdateByID(id, title, author, society, award, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockSocietyController)(nil).UpdateByID), id, title, author, society, award, date)
}

// DeleteByID mocks base method
func (m *MockSocietyController) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockSocietyControllerMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockSocietyController)(nil).DeleteByID), id)
}

// AdminGetAll mocks base method
func (m *MockSocietyController) AdminGetAll() ([]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminGetAll")
	ret0, _ := ret[0].([]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockSocietyControllerMockRecorder) AdminGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockSocietyController)(nil).AdminGetAll))
}

// AdminGetByID mocks base method
func (m *MockSocietyController) AdminGetByID(id int) (*controller.FieldsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminGetByID", id)
	ret0, _ := ret[0].(*controller.FieldsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockSocietyControllerMockRecorder) AdminGetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockSocietyController)(nil).AdminGetByID), id)
}

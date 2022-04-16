// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/UsefulForMe/go-ecommerce/services (interfaces: UserService)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dto "github.com/UsefulForMe/go-ecommerce/dto"
	errs "github.com/UsefulForMe/go-ecommerce/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserService) Create(arg0 dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*dto.CreateUserResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), arg0)
}

// List mocks base method.
func (m *MockUserService) List() (*dto.GetAllUserResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(*dto.GetAllUserResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List))
}

// Login mocks base method.
func (m *MockUserService) Login(arg0 dto.LoginUserRequest) (*dto.LoginUserResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(*dto.LoginUserResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), arg0)
}

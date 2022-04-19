// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/UsefulForMe/go-ecommerce/services (interfaces: FirebaseService)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	auth "firebase.google.com/go/auth"
	errs "github.com/UsefulForMe/go-ecommerce/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockFirebaseService is a mock of FirebaseService interface.
type MockFirebaseService struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseServiceMockRecorder
}

// MockFirebaseServiceMockRecorder is the mock recorder for MockFirebaseService.
type MockFirebaseServiceMockRecorder struct {
	mock *MockFirebaseService
}

// NewMockFirebaseService creates a new mock instance.
func NewMockFirebaseService(ctrl *gomock.Controller) *MockFirebaseService {
	mock := &MockFirebaseService{ctrl: ctrl}
	mock.recorder = &MockFirebaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirebaseService) EXPECT() *MockFirebaseServiceMockRecorder {
	return m.recorder
}

// VerifyIDToken mocks base method.
func (m *MockFirebaseService) VerifyIDToken(arg0 string) (*auth.Token, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", arg0)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockFirebaseServiceMockRecorder) VerifyIDToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockFirebaseService)(nil).VerifyIDToken), arg0)
}
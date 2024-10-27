// Code generated by MockGen. DO NOT EDIT.
// Source: service/delivery_executive_service.go

// Package service is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeliveryExecutiveService is a mock of DeliveryExecutiveService interface.
type MockDeliveryExecutiveService struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryExecutiveServiceMockRecorder
}

// MockDeliveryExecutiveServiceMockRecorder is the mock recorder for MockDeliveryExecutiveService.
type MockDeliveryExecutiveServiceMockRecorder struct {
	mock *MockDeliveryExecutiveService
}

// NewMockDeliveryExecutiveService creates a new mock instance.
func NewMockDeliveryExecutiveService(ctrl *gomock.Controller) *MockDeliveryExecutiveService {
	mock := &MockDeliveryExecutiveService{ctrl: ctrl}
	mock.recorder = &MockDeliveryExecutiveServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeliveryExecutiveService) EXPECT() *MockDeliveryExecutiveServiceMockRecorder {
	return m.recorder
}

// AssignDeliveryExecutive mocks base method.
func (m *MockDeliveryExecutiveService) AssignDeliveryExecutive(orderID, deliveryExecutiveId int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignDeliveryExecutive", orderID, deliveryExecutiveId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignDeliveryExecutive indicates an expected call of AssignDeliveryExecutive.
func (mr *MockDeliveryExecutiveServiceMockRecorder) AssignDeliveryExecutive(orderID, deliveryExecutiveId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignDeliveryExecutive", reflect.TypeOf((*MockDeliveryExecutiveService)(nil).AssignDeliveryExecutive), orderID, deliveryExecutiveId)
}

// CreateDeliveryExecutive mocks base method.
func (m *MockDeliveryExecutiveService) CreateDeliveryExecutive(location string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeliveryExecutive", location)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeliveryExecutive indicates an expected call of CreateDeliveryExecutive.
func (mr *MockDeliveryExecutiveServiceMockRecorder) CreateDeliveryExecutive(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryExecutive", reflect.TypeOf((*MockDeliveryExecutiveService)(nil).CreateDeliveryExecutive), location)
}
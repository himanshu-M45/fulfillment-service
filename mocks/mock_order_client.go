// Code generated by MockGen. DO NOT EDIT.
// Source: clients/order_client.go

// Package clients is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderClient is a mock of OrderClient interface.
type MockOrderClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrderClientMockRecorder
}

// MockOrderClientMockRecorder is the mock recorder for MockOrderClient.
type MockOrderClientMockRecorder struct {
	mock *MockOrderClient
}

// NewMockOrderClient creates a new mock instance.
func NewMockOrderClient(ctrl *gomock.Controller) *MockOrderClient {
	mock := &MockOrderClient{ctrl: ctrl}
	mock.recorder = &MockOrderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderClient) EXPECT() *MockOrderClientMockRecorder {
	return m.recorder
}

// CheckOrderCredibility mocks base method.
func (m *MockOrderClient) CheckOrderCredibility(orderID int32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOrderCredibility", orderID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOrderCredibility indicates an expected call of CheckOrderCredibility.
func (mr *MockOrderClientMockRecorder) CheckOrderCredibility(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOrderCredibility", reflect.TypeOf((*MockOrderClient)(nil).CheckOrderCredibility), orderID)
}

// UpdateOrderStatus mocks base method.
func (m *MockOrderClient) UpdateOrderStatus(orderID int32, status string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", orderID, status)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockOrderClientMockRecorder) UpdateOrderStatus(orderID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockOrderClient)(nil).UpdateOrderStatus), orderID, status)
}

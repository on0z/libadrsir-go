// Code generated by MockGen. DO NOT EDIT.
// Source: ./libadrsir.go

// Package mock_libadrsir is a generated GoMock package.
package mock_libadrsir

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBus is a mock of Bus interface.
type MockBus struct {
	ctrl     *gomock.Controller
	recorder *MockBusMockRecorder
}

// MockBusMockRecorder is the mock recorder for MockBus.
type MockBusMockRecorder struct {
	mock *MockBus
}

// NewMockBus creates a new mock instance.
func NewMockBus(ctrl *gomock.Controller) *MockBus {
	mock := &MockBus{ctrl: ctrl}
	mock.recorder = &MockBusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBus) EXPECT() *MockBusMockRecorder {
	return m.recorder
}

// Tx mocks base method.
func (m *MockBus) Tx(w, r []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", w, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tx indicates an expected call of Tx.
func (mr *MockBusMockRecorder) Tx(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockBus)(nil).Tx), w, r)
}

// MockAdrsirAPI is a mock of AdrsirAPI interface.
type MockAdrsirAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAdrsirAPIMockRecorder
}

// MockAdrsirAPIMockRecorder is the mock recorder for MockAdrsirAPI.
type MockAdrsirAPIMockRecorder struct {
	mock *MockAdrsirAPI
}

// NewMockAdrsirAPI creates a new mock instance.
func NewMockAdrsirAPI(ctrl *gomock.Controller) *MockAdrsirAPI {
	mock := &MockAdrsirAPI{ctrl: ctrl}
	mock.recorder = &MockAdrsirAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdrsirAPI) EXPECT() *MockAdrsirAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAdrsirAPI) Get(index int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", index)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockAdrsirAPIMockRecorder) Get(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAdrsirAPI)(nil).Get), index)
}

// Send mocks base method.
func (m *MockAdrsirAPI) Send(irCommandStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", irCommandStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAdrsirAPIMockRecorder) Send(irCommandStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAdrsirAPI)(nil).Send), irCommandStr)
}

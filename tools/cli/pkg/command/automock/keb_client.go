// Code generated by MockGen. DO NOT EDIT.
// Source: reconciliations.go

// Package automock is a generated GoMock package.
package automock

import (
	gomock "github.com/golang/mock/gomock"
	runtime "github.com/kyma-project/control-plane/components/kyma-environment-broker/common/runtime"
	reflect "reflect"
)

// MockkebClient is a mock of kebClient interface
type MockkebClient struct {
	ctrl     *gomock.Controller
	recorder *MockkebClientMockRecorder
}

// MockkebClientMockRecorder is the mock recorder for MockkebClient
type MockkebClientMockRecorder struct {
	mock *MockkebClient
}

// NewMockkebClient creates a new mock instance
func NewMockkebClient(ctrl *gomock.Controller) *MockkebClient {
	mock := &MockkebClient{ctrl: ctrl}
	mock.recorder = &MockkebClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockkebClient) EXPECT() *MockkebClientMockRecorder {
	return m.recorder
}

// ListRuntimes mocks base method
func (m *MockkebClient) ListRuntimes(arg0 runtime.ListParameters) (runtime.RuntimesPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRuntimes", arg0)
	ret0, _ := ret[0].(runtime.RuntimesPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuntimes indicates an expected call of ListRuntimes
func (mr *MockkebClientMockRecorder) ListRuntimes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuntimes", reflect.TypeOf((*MockkebClient)(nil).ListRuntimes), arg0)
}
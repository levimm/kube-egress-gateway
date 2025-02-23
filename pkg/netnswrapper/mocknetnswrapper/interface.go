// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/netnswrapper/netns_linux.go

// Package mocknetnswrapper is a generated GoMock package.
package mocknetnswrapper

import (
	reflect "reflect"

	ns "github.com/containernetworking/plugins/pkg/ns"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// GetNS mocks base method.
func (m *MockInterface) GetNS(nsName string) (ns.NetNS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNS", nsName)
	ret0, _ := ret[0].(ns.NetNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNS indicates an expected call of GetNS.
func (mr *MockInterfaceMockRecorder) GetNS(nsName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNS", reflect.TypeOf((*MockInterface)(nil).GetNS), nsName)
}

// GetNSByPath mocks base method.
func (m *MockInterface) GetNSByPath(nsPath string) (ns.NetNS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNSByPath", nsPath)
	ret0, _ := ret[0].(ns.NetNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNSByPath indicates an expected call of GetNSByPath.
func (mr *MockInterfaceMockRecorder) GetNSByPath(nsPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNSByPath", reflect.TypeOf((*MockInterface)(nil).GetNSByPath), nsPath)
}

// ListNS mocks base method.
func (m *MockInterface) ListNS() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNS")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNS indicates an expected call of ListNS.
func (mr *MockInterfaceMockRecorder) ListNS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNS", reflect.TypeOf((*MockInterface)(nil).ListNS))
}

// NewNS mocks base method.
func (m *MockInterface) NewNS(nsName string) (ns.NetNS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNS", nsName)
	ret0, _ := ret[0].(ns.NetNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNS indicates an expected call of NewNS.
func (mr *MockInterfaceMockRecorder) NewNS(nsName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNS", reflect.TypeOf((*MockInterface)(nil).NewNS), nsName)
}

// UnmountNS mocks base method.
func (m *MockInterface) UnmountNS(nsName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmountNS", nsName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmountNS indicates an expected call of UnmountNS.
func (mr *MockInterfaceMockRecorder) UnmountNS(nsName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmountNS", reflect.TypeOf((*MockInterface)(nil).UnmountNS), nsName)
}

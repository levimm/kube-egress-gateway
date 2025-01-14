// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/netlinkwrapper/netlink.go

// Package mocknetlinkwrapper is a generated GoMock package.
package mocknetlinkwrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
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

// AddrAdd mocks base method.
func (m *MockInterface) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrAdd", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd.
func (mr *MockInterfaceMockRecorder) AddrAdd(link, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockInterface)(nil).AddrAdd), link, addr)
}

// AddrDel mocks base method.
func (m *MockInterface) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrDel", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrDel indicates an expected call of AddrDel.
func (mr *MockInterfaceMockRecorder) AddrDel(link, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrDel", reflect.TypeOf((*MockInterface)(nil).AddrDel), link, addr)
}

// AddrList mocks base method.
func (m *MockInterface) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrList", link, family)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddrList indicates an expected call of AddrList.
func (mr *MockInterfaceMockRecorder) AddrList(link, family interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockInterface)(nil).AddrList), link, family)
}

// AddrReplace mocks base method.
func (m *MockInterface) AddrReplace(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrReplace", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrReplace indicates an expected call of AddrReplace.
func (mr *MockInterfaceMockRecorder) AddrReplace(link, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrReplace", reflect.TypeOf((*MockInterface)(nil).AddrReplace), link, addr)
}

// LinkAdd mocks base method.
func (m *MockInterface) LinkAdd(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkAdd", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkAdd indicates an expected call of LinkAdd.
func (mr *MockInterfaceMockRecorder) LinkAdd(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkAdd", reflect.TypeOf((*MockInterface)(nil).LinkAdd), link)
}

// LinkByName mocks base method.
func (m *MockInterface) LinkByName(name string) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByName", name)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByName indicates an expected call of LinkByName.
func (mr *MockInterfaceMockRecorder) LinkByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByName", reflect.TypeOf((*MockInterface)(nil).LinkByName), name)
}

// LinkDel mocks base method.
func (m *MockInterface) LinkDel(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkDel", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDel indicates an expected call of LinkDel.
func (mr *MockInterfaceMockRecorder) LinkDel(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDel", reflect.TypeOf((*MockInterface)(nil).LinkDel), link)
}

// LinkSetAlias mocks base method.
func (m *MockInterface) LinkSetAlias(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetAlias", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetAlias indicates an expected call of LinkSetAlias.
func (mr *MockInterfaceMockRecorder) LinkSetAlias(link, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetAlias", reflect.TypeOf((*MockInterface)(nil).LinkSetAlias), link, name)
}

// LinkSetDown mocks base method.
func (m *MockInterface) LinkSetDown(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetDown", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetDown indicates an expected call of LinkSetDown.
func (mr *MockInterfaceMockRecorder) LinkSetDown(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetDown", reflect.TypeOf((*MockInterface)(nil).LinkSetDown), link)
}

// LinkSetName mocks base method.
func (m *MockInterface) LinkSetName(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetName", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetName indicates an expected call of LinkSetName.
func (mr *MockInterfaceMockRecorder) LinkSetName(link, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetName", reflect.TypeOf((*MockInterface)(nil).LinkSetName), link, name)
}

// LinkSetNsFd mocks base method.
func (m *MockInterface) LinkSetNsFd(link netlink.Link, fd int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetNsFd", link, fd)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetNsFd indicates an expected call of LinkSetNsFd.
func (mr *MockInterfaceMockRecorder) LinkSetNsFd(link, fd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetNsFd", reflect.TypeOf((*MockInterface)(nil).LinkSetNsFd), link, fd)
}

// LinkSetUp mocks base method.
func (m *MockInterface) LinkSetUp(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetUp", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp.
func (mr *MockInterfaceMockRecorder) LinkSetUp(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockInterface)(nil).LinkSetUp), link)
}

// RouteDel mocks base method.
func (m *MockInterface) RouteDel(route *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteDel", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteDel indicates an expected call of RouteDel.
func (mr *MockInterfaceMockRecorder) RouteDel(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDel", reflect.TypeOf((*MockInterface)(nil).RouteDel), route)
}

// RouteList mocks base method.
func (m *MockInterface) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteList", link, family)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteList indicates an expected call of RouteList.
func (mr *MockInterfaceMockRecorder) RouteList(link, family interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteList", reflect.TypeOf((*MockInterface)(nil).RouteList), link, family)
}

// RouteReplace mocks base method.
func (m *MockInterface) RouteReplace(route *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteReplace", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteReplace indicates an expected call of RouteReplace.
func (mr *MockInterfaceMockRecorder) RouteReplace(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteReplace", reflect.TypeOf((*MockInterface)(nil).RouteReplace), route)
}

// RuleAdd mocks base method.
func (m *MockInterface) RuleAdd(rule *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleAdd", rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd.
func (mr *MockInterfaceMockRecorder) RuleAdd(rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockInterface)(nil).RuleAdd), rule)
}

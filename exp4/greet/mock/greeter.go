// Code generated by MockGen. DO NOT EDIT.
// Source: greet/greeter.go

// Package mock_greet is a generated GoMock package.
package mock_greet

import (
	greet "party/greet"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGreeter is a mock of Greeter interface.
type MockGreeter struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterMockRecorder
}

// MockGreeterMockRecorder is the mock recorder for MockGreeter.
type MockGreeterMockRecorder struct {
	mock *MockGreeter
}

// NewMockGreeter creates a new mock instance.
func NewMockGreeter(ctrl *gomock.Controller) *MockGreeter {
	mock := &MockGreeter{ctrl: ctrl}
	mock.recorder = &MockGreeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeter) EXPECT() *MockGreeterMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockGreeter) Hello(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Hello indicates an expected call of Hello.
func (mr *MockGreeterMockRecorder) Hello(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockGreeter)(nil).Hello), name)
}

// Hello1 mocks base method.
func (m *MockGreeter) Hello1(name1, name2 string) (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello1", name1, name2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// Hello1 indicates an expected call of Hello1.
func (mr *MockGreeterMockRecorder) Hello1(name1, name2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello1", reflect.TypeOf((*MockGreeter)(nil).Hello1), name1, name2)
}

// MockVisitorLister is a mock of VisitorLister interface.
type MockVisitorLister struct {
	ctrl     *gomock.Controller
	recorder *MockVisitorListerMockRecorder
}

// MockVisitorListerMockRecorder is the mock recorder for MockVisitorLister.
type MockVisitorListerMockRecorder struct {
	mock *MockVisitorLister
}

// NewMockVisitorLister creates a new mock instance.
func NewMockVisitorLister(ctrl *gomock.Controller) *MockVisitorLister {
	mock := &MockVisitorLister{ctrl: ctrl}
	mock.recorder = &MockVisitorListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisitorLister) EXPECT() *MockVisitorListerMockRecorder {
	return m.recorder
}

// ListVisitors mocks base method.
func (m *MockVisitorLister) ListVisitors(who greet.VisitorGroup) ([]greet.Visitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVisitors", who)
	ret0, _ := ret[0].([]greet.Visitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVisitors indicates an expected call of ListVisitors.
func (mr *MockVisitorListerMockRecorder) ListVisitors(who interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVisitors", reflect.TypeOf((*MockVisitorLister)(nil).ListVisitors), who)
}

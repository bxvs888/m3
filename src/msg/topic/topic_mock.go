// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/msg/topic (interfaces: Service)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package topic is a generated GoMock package.
package topic

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CheckAndSet mocks base method
func (m *MockService) CheckAndSet(arg0 Topic, arg1 int) (Topic, error) {
	ret := m.ctrl.Call(m, "CheckAndSet", arg0, arg1)
	ret0, _ := ret[0].(Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndSet indicates an expected call of CheckAndSet
func (mr *MockServiceMockRecorder) CheckAndSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndSet", reflect.TypeOf((*MockService)(nil).CheckAndSet), arg0, arg1)
}

// Delete mocks base method
func (m *MockService) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockService) Get(arg0 string) (Topic, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockService)(nil).Get), arg0)
}

// Watch mocks base method
func (m *MockService) Watch(arg0 string) (Watch, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockServiceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockService)(nil).Watch), arg0)
}
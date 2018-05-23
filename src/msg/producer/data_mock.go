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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3msg/producer (interfaces: Data)

package producer

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockData is a mock of Data interface
type MockData struct {
	ctrl     *gomock.Controller
	recorder *MockDataMockRecorder
}

// MockDataMockRecorder is the mock recorder for MockData
type MockDataMockRecorder struct {
	mock *MockData
}

// NewMockData creates a new mock instance
func NewMockData(ctrl *gomock.Controller) *MockData {
	mock := &MockData{ctrl: ctrl}
	mock.recorder = &MockDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockData) EXPECT() *MockDataMockRecorder {
	return _m.recorder
}

// Bytes mocks base method
func (_m *MockData) Bytes() []byte {
	ret := _m.ctrl.Call(_m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes
func (_mr *MockDataMockRecorder) Bytes() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Bytes", reflect.TypeOf((*MockData)(nil).Bytes))
}

// Finalize mocks base method
func (_m *MockData) Finalize(_param0 FinalizeReason) {
	_m.ctrl.Call(_m, "Finalize", _param0)
}

// Finalize indicates an expected call of Finalize
func (_mr *MockDataMockRecorder) Finalize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Finalize", reflect.TypeOf((*MockData)(nil).Finalize), arg0)
}

// Shard mocks base method
func (_m *MockData) Shard() uint32 {
	ret := _m.ctrl.Call(_m, "Shard")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Shard indicates an expected call of Shard
func (_mr *MockDataMockRecorder) Shard() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Shard", reflect.TypeOf((*MockData)(nil).Shard))
}

// Size mocks base method
func (_m *MockData) Size() uint32 {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockDataMockRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Size", reflect.TypeOf((*MockData)(nil).Size))
}
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3aggregator/aggregator (interfaces: ElectionManager,FlushTimesManager,PlacementManager)

package aggregator

import (
	"context"

	"github.com/m3db/m3aggregator/generated/proto/flush"
	"github.com/m3db/m3cluster/placement"
	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3x/watch"

	"github.com/golang/mock/gomock"
)

// Mock of ElectionManager interface
type MockElectionManager struct {
	ctrl     *gomock.Controller
	recorder *_MockElectionManagerRecorder
}

// Recorder for MockElectionManager (not exported)
type _MockElectionManagerRecorder struct {
	mock *MockElectionManager
}

func NewMockElectionManager(ctrl *gomock.Controller) *MockElectionManager {
	mock := &MockElectionManager{ctrl: ctrl}
	mock.recorder = &_MockElectionManagerRecorder{mock}
	return mock
}

func (_m *MockElectionManager) EXPECT() *_MockElectionManagerRecorder {
	return _m.recorder
}

func (_m *MockElectionManager) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockElectionManager) ElectionState() ElectionState {
	ret := _m.ctrl.Call(_m, "ElectionState")
	ret0, _ := ret[0].(ElectionState)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) ElectionState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ElectionState")
}

func (_m *MockElectionManager) IsCampaigning() bool {
	ret := _m.ctrl.Call(_m, "IsCampaigning")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) IsCampaigning() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsCampaigning")
}

func (_m *MockElectionManager) Open(_param0 uint32) error {
	ret := _m.ctrl.Call(_m, "Open", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}

func (_m *MockElectionManager) Reset() error {
	ret := _m.ctrl.Call(_m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockElectionManager) Resign(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Resign", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockElectionManagerRecorder) Resign(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resign", arg0)
}

// Mock of FlushTimesManager interface
type MockFlushTimesManager struct {
	ctrl     *gomock.Controller
	recorder *_MockFlushTimesManagerRecorder
}

// Recorder for MockFlushTimesManager (not exported)
type _MockFlushTimesManagerRecorder struct {
	mock *MockFlushTimesManager
}

func NewMockFlushTimesManager(ctrl *gomock.Controller) *MockFlushTimesManager {
	mock := &MockFlushTimesManager{ctrl: ctrl}
	mock.recorder = &_MockFlushTimesManagerRecorder{mock}
	return mock
}

func (_m *MockFlushTimesManager) EXPECT() *_MockFlushTimesManagerRecorder {
	return _m.recorder
}

func (_m *MockFlushTimesManager) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFlushTimesManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockFlushTimesManager) Get() (*flush.ShardSetFlushTimes, error) {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(*flush.ShardSetFlushTimes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFlushTimesManagerRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockFlushTimesManager) Open(_param0 uint32) error {
	ret := _m.ctrl.Call(_m, "Open", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFlushTimesManagerRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}

func (_m *MockFlushTimesManager) Reset() error {
	ret := _m.ctrl.Call(_m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFlushTimesManagerRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockFlushTimesManager) StoreAsync(_param0 *flush.ShardSetFlushTimes) error {
	ret := _m.ctrl.Call(_m, "StoreAsync", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFlushTimesManagerRecorder) StoreAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StoreAsync", arg0)
}

func (_m *MockFlushTimesManager) Watch() (watch.Watch, error) {
	ret := _m.ctrl.Call(_m, "Watch")
	ret0, _ := ret[0].(watch.Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFlushTimesManagerRecorder) Watch() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Watch")
}

// Mock of PlacementManager interface
type MockPlacementManager struct {
	ctrl     *gomock.Controller
	recorder *_MockPlacementManagerRecorder
}

// Recorder for MockPlacementManager (not exported)
type _MockPlacementManagerRecorder struct {
	mock *MockPlacementManager
}

func NewMockPlacementManager(ctrl *gomock.Controller) *MockPlacementManager {
	mock := &MockPlacementManager{ctrl: ctrl}
	mock.recorder = &_MockPlacementManagerRecorder{mock}
	return mock
}

func (_m *MockPlacementManager) EXPECT() *_MockPlacementManagerRecorder {
	return _m.recorder
}

func (_m *MockPlacementManager) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPlacementManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockPlacementManager) HasReplacementInstance() (bool, error) {
	ret := _m.ctrl.Call(_m, "HasReplacementInstance")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPlacementManagerRecorder) HasReplacementInstance() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasReplacementInstance")
}

func (_m *MockPlacementManager) Instance() (placement.Instance, error) {
	ret := _m.ctrl.Call(_m, "Instance")
	ret0, _ := ret[0].(placement.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPlacementManagerRecorder) Instance() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Instance")
}

func (_m *MockPlacementManager) InstanceFrom(_param0 placement.Placement) (placement.Instance, error) {
	ret := _m.ctrl.Call(_m, "InstanceFrom", _param0)
	ret0, _ := ret[0].(placement.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPlacementManagerRecorder) InstanceFrom(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstanceFrom", arg0)
}

func (_m *MockPlacementManager) Open() error {
	ret := _m.ctrl.Call(_m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPlacementManagerRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

func (_m *MockPlacementManager) Placement() (placement.ActiveStagedPlacement, placement.Placement, error) {
	ret := _m.ctrl.Call(_m, "Placement")
	ret0, _ := ret[0].(placement.ActiveStagedPlacement)
	ret1, _ := ret[1].(placement.Placement)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPlacementManagerRecorder) Placement() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Placement")
}

func (_m *MockPlacementManager) Shards() (shard.Shards, error) {
	ret := _m.ctrl.Call(_m, "Shards")
	ret0, _ := ret[0].(shard.Shards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPlacementManagerRecorder) Shards() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shards")
}
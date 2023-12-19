// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen --source interfaces.go --destination mock_eventhandler/mock_handlers.go --package mock_eventhandler
//
// Package mock_eventhandler is a generated GoMock package.
package mock_eventhandler

import (
	context "context"
	reflect "reflect"

	model "github.com/bacalhau-project/bacalhau/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockJobEventHandler is a mock of JobEventHandler interface.
type MockJobEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJobEventHandlerMockRecorder
}

// MockJobEventHandlerMockRecorder is the mock recorder for MockJobEventHandler.
type MockJobEventHandlerMockRecorder struct {
	mock *MockJobEventHandler
}

// NewMockJobEventHandler creates a new mock instance.
func NewMockJobEventHandler(ctrl *gomock.Controller) *MockJobEventHandler {
	mock := &MockJobEventHandler{ctrl: ctrl}
	mock.recorder = &MockJobEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobEventHandler) EXPECT() *MockJobEventHandlerMockRecorder {
	return m.recorder
}

// HandleJobEvent mocks base method.
func (m *MockJobEventHandler) HandleJobEvent(ctx context.Context, event model.JobEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleJobEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleJobEvent indicates an expected call of HandleJobEvent.
func (mr *MockJobEventHandlerMockRecorder) HandleJobEvent(ctx, event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleJobEvent", reflect.TypeOf((*MockJobEventHandler)(nil).HandleJobEvent), ctx, event)
}

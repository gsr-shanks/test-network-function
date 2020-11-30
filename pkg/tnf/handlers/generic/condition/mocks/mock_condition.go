// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/tnf/handlers/generic/condition/condition.go

// Package mock_condition is a generated GoMock package.
package mock_condition

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	regexp "regexp"
)

// MockCondition is a mock of Condition interface
type MockCondition struct {
	ctrl     *gomock.Controller
	recorder *MockConditionMockRecorder
}

// MockConditionMockRecorder is the mock recorder for MockCondition
type MockConditionMockRecorder struct {
	mock *MockCondition
}

// NewMockCondition creates a new mock instance
func NewMockCondition(ctrl *gomock.Controller) *MockCondition {
	mock := &MockCondition{ctrl: ctrl}
	mock.recorder = &MockConditionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCondition) EXPECT() *MockConditionMockRecorder {
	return m.recorder
}

// Evaluate mocks base method
func (m *MockCondition) Evaluate(match string, regex regexp.Regexp, groupIdx int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Evaluate", match, regex, groupIdx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Evaluate indicates an expected call of Evaluate
func (mr *MockConditionMockRecorder) Evaluate(match, regex, groupIdx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evaluate", reflect.TypeOf((*MockCondition)(nil).Evaluate), match, regex, groupIdx)
}
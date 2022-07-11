// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	mock "github.com/stretchr/testify/mock"

	sync "sync"
)

// AfterExtraction is an autogenerated mock type for the AfterExtraction type
type AfterExtraction struct {
	mock.Mock
}

// Execute provides a mock function with given fields: tykTaskConfigs
func (_m *AfterExtraction) Execute(tykTaskConfigs []model.TykTaskConfig) {
	_m.Called(tykTaskConfigs)
}

// ExecuteAsync provides a mock function with given fields: tykTaskConfigs, wg
func (_m *AfterExtraction) ExecuteAsync(tykTaskConfigs []model.TykTaskConfig, wg *sync.WaitGroup) {
	_m.Called(tykTaskConfigs, wg)
}

type mockConstructorTestingTNewAfterExtraction interface {
	mock.TestingT
	Cleanup(func())
}

// NewAfterExtraction creates a new instance of AfterExtraction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAfterExtraction(t mockConstructorTestingTNewAfterExtraction) *AfterExtraction {
	mock := &AfterExtraction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
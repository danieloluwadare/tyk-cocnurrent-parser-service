// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	mock "github.com/stretchr/testify/mock"
)

// Parser is an autogenerated mock type for the Parser type
type Parser struct {
	mock.Mock
}

// Parse provides a mock function with given fields:
func (_m *Parser) Parse() ([]model.TykTaskConfig, error) {
	ret := _m.Called()

	var r0 []model.TykTaskConfig
	if rf, ok := ret.Get(0).(func() []model.TykTaskConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TykTaskConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewParser interface {
	mock.TestingT
	Cleanup(func())
}

// NewParser creates a new instance of Parser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewParser(t mockConstructorTestingTNewParser) *Parser {
	mock := &Parser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

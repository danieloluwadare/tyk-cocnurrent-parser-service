// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	mock "github.com/stretchr/testify/mock"
)

// InMemoryServiceServer is an autogenerated mock type for the InMemoryServiceServer type
type InMemoryServiceServer struct {
	mock.Mock
}

// SaveRequest provides a mock function with given fields: _a0, _a1
func (_m *InMemoryServiceServer) SaveRequest(_a0 context.Context, _a1 *proto.TykServerRequest) (*proto.TykServerResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.TykServerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.TykServerRequest) *proto.TykServerResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.TykServerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.TykServerRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedInMemoryServiceServer provides a mock function with given fields:
func (_m *InMemoryServiceServer) mustEmbedUnimplementedInMemoryServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewInMemoryServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewInMemoryServiceServer creates a new instance of InMemoryServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInMemoryServiceServer(t mockConstructorTestingTNewInMemoryServiceServer) *InMemoryServiceServer {
	mock := &InMemoryServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

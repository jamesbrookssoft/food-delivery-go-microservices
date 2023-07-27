// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	consumer "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/consumer"
	mock "github.com/stretchr/testify/mock"

	types "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/types"
)

// ConsumerConnector is an autogenerated mock type for the ConsumerConnector type
type ConsumerConnector struct {
	mock.Mock
}

// ConnectConsumer provides a mock function with given fields: messageType, _a1
func (_m *ConsumerConnector) ConnectConsumer(messageType types.IMessage, _a1 consumer.Consumer) error {
	ret := _m.Called(messageType, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.IMessage, consumer.Consumer) error); ok {
		r0 = rf(messageType, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectConsumerHandler provides a mock function with given fields: messageType, consumerHandler
func (_m *ConsumerConnector) ConnectConsumerHandler(messageType types.IMessage, consumerHandler consumer.ConsumerHandler) error {
	ret := _m.Called(messageType, consumerHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.IMessage, consumer.ConsumerHandler) error); ok {
		r0 = rf(messageType, consumerHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConsumerConnector interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerConnector creates a new instance of ConsumerConnector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerConnector(t mockConstructorTestingTNewConsumerConnector) *ConsumerConnector {
	mock := &ConsumerConnector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	consumer "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/consumer"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerHandlerConfigurationBuilder is an autogenerated mock type for the ConsumerHandlerConfigurationBuilder type
type ConsumerHandlerConfigurationBuilder struct {
	mock.Mock
}

// AddHandler provides a mock function with given fields: handler
func (_m *ConsumerHandlerConfigurationBuilder) AddHandler(handler consumer.ConsumerHandler) consumer.ConsumerHandlerConfigurationBuilder {
	ret := _m.Called(handler)

	var r0 consumer.ConsumerHandlerConfigurationBuilder
	if rf, ok := ret.Get(0).(func(consumer.ConsumerHandler) consumer.ConsumerHandlerConfigurationBuilder); ok {
		r0 = rf(handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(consumer.ConsumerHandlerConfigurationBuilder)
		}
	}

	return r0
}

// Build provides a mock function with given fields:
func (_m *ConsumerHandlerConfigurationBuilder) Build() *consumer.ConsumerHandlersConfiguration {
	ret := _m.Called()

	var r0 *consumer.ConsumerHandlersConfiguration
	if rf, ok := ret.Get(0).(func() *consumer.ConsumerHandlersConfiguration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consumer.ConsumerHandlersConfiguration)
		}
	}

	return r0
}

type mockConstructorTestingTNewConsumerHandlerConfigurationBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerHandlerConfigurationBuilder creates a new instance of ConsumerHandlerConfigurationBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerHandlerConfigurationBuilder(t mockConstructorTestingTNewConsumerHandlerConfigurationBuilder) *ConsumerHandlerConfigurationBuilder {
	mock := &ConsumerHandlerConfigurationBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

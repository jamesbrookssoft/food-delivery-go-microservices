// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/core/domain"
	metadata "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/core/metadata"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/satori/go.uuid"
)

// IHaveEventSourcedAggregate is an autogenerated mock type for the IHaveEventSourcedAggregate type
type IHaveEventSourcedAggregate struct {
	mock.Mock
}

// AddDomainEvents provides a mock function with given fields: event
func (_m *IHaveEventSourcedAggregate) AddDomainEvents(event domain.IDomainEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Apply provides a mock function with given fields: event, isNew
func (_m *IHaveEventSourcedAggregate) Apply(event domain.IDomainEvent, isNew bool) error {
	ret := _m.Called(event, isNew)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent, bool) error); ok {
		r0 = rf(event, isNew)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatedAt provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) CreatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// CurrentVersion provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) CurrentVersion() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// HasUncommittedEvents provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) HasUncommittedEvents() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Id provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) Id() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// LoadFromHistory provides a mock function with given fields: events, _a1
func (_m *IHaveEventSourcedAggregate) LoadFromHistory(events []domain.IDomainEvent, _a1 metadata.Metadata) error {
	ret := _m.Called(events, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.IDomainEvent, metadata.Metadata) error); ok {
		r0 = rf(events, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkUncommittedEventAsCommitted provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) MarkUncommittedEventAsCommitted() {
	_m.Called()
}

// NewEmptyAggregate provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) NewEmptyAggregate() {
	_m.Called()
}

// OriginalVersion provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) OriginalVersion() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// SetEntityType provides a mock function with given fields: entityType
func (_m *IHaveEventSourcedAggregate) SetEntityType(entityType string) {
	_m.Called(entityType)
}

// SetId provides a mock function with given fields: id
func (_m *IHaveEventSourcedAggregate) SetId(id uuid.UUID) {
	_m.Called(id)
}

// SetOriginalVersion provides a mock function with given fields: version
func (_m *IHaveEventSourcedAggregate) SetOriginalVersion(version int64) {
	_m.Called(version)
}

// SetUpdatedAt provides a mock function with given fields: updatedAt
func (_m *IHaveEventSourcedAggregate) SetUpdatedAt(updatedAt time.Time) {
	_m.Called(updatedAt)
}

// UncommittedEvents provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) UncommittedEvents() []domain.IDomainEvent {
	ret := _m.Called()

	var r0 []domain.IDomainEvent
	if rf, ok := ret.Get(0).(func() []domain.IDomainEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IDomainEvent)
		}
	}

	return r0
}

// UpdatedAt provides a mock function with given fields:
func (_m *IHaveEventSourcedAggregate) UpdatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// When provides a mock function with given fields: event
func (_m *IHaveEventSourcedAggregate) When(event domain.IDomainEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// fold provides a mock function with given fields: event, _a1
func (_m *IHaveEventSourcedAggregate) fold(event domain.IDomainEvent, _a1 metadata.Metadata) error {
	ret := _m.Called(event, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent, metadata.Metadata) error); ok {
		r0 = rf(event, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIHaveEventSourcedAggregate interface {
	mock.TestingT
	Cleanup(func())
}

// NewIHaveEventSourcedAggregate creates a new instance of IHaveEventSourcedAggregate. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIHaveEventSourcedAggregate(t mockConstructorTestingTNewIHaveEventSourcedAggregate) *IHaveEventSourcedAggregate {
	mock := &IHaveEventSourcedAggregate{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

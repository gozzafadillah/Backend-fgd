// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	topics "charum/business/topics"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CreateTopic provides a mock function with given fields: domain
func (_m *UseCase) CreateTopic(domain *topics.Domain) (topics.Domain, error) {
	ret := _m.Called(domain)

	var r0 topics.Domain
	if rf, ok := ret.Get(0).(func(*topics.Domain) topics.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(topics.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*topics.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTopic provides a mock function with given fields: id
func (_m *UseCase) DeleteTopic(id primitive.ObjectID) (topics.Domain, error) {
	ret := _m.Called(id)

	var r0 topics.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) topics.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(topics.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *UseCase) GetByID(id primitive.ObjectID) (topics.Domain, error) {
	ret := _m.Called(id)

	var r0 topics.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) topics.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(topics.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTopic provides a mock function with given fields: id, domain
func (_m *UseCase) UpdateTopic(id primitive.ObjectID, domain *topics.Domain) (topics.Domain, error) {
	ret := _m.Called(id, domain)

	var r0 topics.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *topics.Domain) topics.Domain); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Get(0).(topics.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, *topics.Domain) error); ok {
		r1 = rf(id, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

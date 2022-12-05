// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	follow_threads "charum/business/follow_threads"
	dtofollow_threads "charum/dto/follow_threads"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CountByThreadID provides a mock function with given fields: threadID
func (_m *UseCase) CountByThreadID(threadID primitive.ObjectID) (int, error) {
	ret := _m.Called(threadID)

	var r0 int
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) int); ok {
		r0 = rf(threadID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(threadID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: domain
func (_m *UseCase) Create(domain *follow_threads.Domain) (follow_threads.Domain, error) {
	ret := _m.Called(domain)

	var r0 follow_threads.Domain
	if rf, ok := ret.Get(0).(func(*follow_threads.Domain) follow_threads.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(follow_threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*follow_threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: domain
func (_m *UseCase) Delete(domain *follow_threads.Domain) (follow_threads.Domain, error) {
	ret := _m.Called(domain)

	var r0 follow_threads.Domain
	if rf, ok := ret.Get(0).(func(*follow_threads.Domain) follow_threads.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(follow_threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*follow_threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllByThreadID provides a mock function with given fields: threadID
func (_m *UseCase) DeleteAllByThreadID(threadID primitive.ObjectID) error {
	ret := _m.Called(threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByUserID provides a mock function with given fields: id
func (_m *UseCase) DeleteAllByUserID(id primitive.ObjectID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DomainToResponse provides a mock function with given fields: domain
func (_m *UseCase) DomainToResponse(domain follow_threads.Domain) (dtofollow_threads.Response, error) {
	ret := _m.Called(domain)

	var r0 dtofollow_threads.Response
	if rf, ok := ret.Get(0).(func(follow_threads.Domain) dtofollow_threads.Response); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(dtofollow_threads.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(follow_threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DomainToResponseArray provides a mock function with given fields: domain
func (_m *UseCase) DomainToResponseArray(domain []follow_threads.Domain) ([]dtofollow_threads.Response, error) {
	ret := _m.Called(domain)

	var r0 []dtofollow_threads.Response
	if rf, ok := ret.Get(0).(func([]follow_threads.Domain) []dtofollow_threads.Response); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtofollow_threads.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]follow_threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllByUserID provides a mock function with given fields: userID
func (_m *UseCase) GetAllByUserID(userID primitive.ObjectID) ([]follow_threads.Domain, error) {
	ret := _m.Called(userID)

	var r0 []follow_threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) []follow_threads.Domain); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]follow_threads.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetNotification provides a mock function with given fields: threadID, userID
func (_m *UseCase) ResetNotification(threadID primitive.ObjectID, userID primitive.ObjectID) error {
	ret := _m.Called(threadID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(threadID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotification provides a mock function with given fields: threadID
func (_m *UseCase) UpdateNotification(threadID primitive.ObjectID) error {
	ret := _m.Called(threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

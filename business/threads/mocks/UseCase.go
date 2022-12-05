// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "charum/dto"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	threads "charum/business/threads"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: domain
func (_m *UseCase) Create(domain *threads.Domain) (threads.Domain, error) {
	ret := _m.Called(domain)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(*threads.Domain) threads.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, threadID
func (_m *UseCase) Delete(userID primitive.ObjectID, threadID primitive.ObjectID) (threads.Domain, error) {
	ret := _m.Called(userID, threadID)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) threads.Domain); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r1 = rf(userID, threadID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DomainToResponse provides a mock function with given fields: domain
func (_m *UseCase) DomainToResponse(domain threads.Domain) (dto.ResponseThread, error) {
	ret := _m.Called(domain)

	var r0 dto.ResponseThread
	if rf, ok := ret.Get(0).(func(threads.Domain) dto.ResponseThread); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(dto.ResponseThread)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(threads.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DomainsToResponseArray provides a mock function with given fields: domains
func (_m *UseCase) DomainsToResponseArray(domains []threads.Domain) ([]dto.ResponseThread, error) {
	ret := _m.Called(domains)

	var r0 []dto.ResponseThread
	if rf, ok := ret.Get(0).(func([]threads.Domain) []dto.ResponseThread); ok {
		r0 = rf(domains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ResponseThread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]threads.Domain) error); ok {
		r1 = rf(domains)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *UseCase) GetByID(id primitive.ObjectID) (threads.Domain, error) {
	ret := _m.Called(id)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) threads.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithSortAndOrder provides a mock function with given fields: page, limit, sort, order
func (_m *UseCase) GetWithSortAndOrder(page int, limit int, sort string, order string) ([]threads.Domain, int, int, error) {
	ret := _m.Called(page, limit, sort, order)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(int, int, string, string) []threads.Domain); ok {
		r0 = rf(page, limit, sort, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, string, string) int); ok {
		r1 = rf(page, limit, sort, order)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(int, int, string, string) int); ok {
		r2 = rf(page, limit, sort, order)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(int, int, string, string) error); ok {
		r3 = rf(page, limit, sort, order)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SuspendByUserID provides a mock function with given fields: userID
func (_m *UseCase) SuspendByUserID(userID primitive.ObjectID) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: domain
func (_m *UseCase) Update(domain *threads.Domain) (threads.Domain, error) {
	ret := _m.Called(domain)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(*threads.Domain) threads.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*threads.Domain) error); ok {
		r1 = rf(domain)
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

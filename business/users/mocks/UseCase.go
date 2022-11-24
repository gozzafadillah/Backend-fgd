// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	users "charum/business/users"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UseCase) Delete(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
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
func (_m *UseCase) GetByID(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
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
func (_m *UseCase) GetWithSortAndOrder(page int, limit int, sort string, order string) ([]users.Domain, int, error) {
	ret := _m.Called(page, limit, sort, order)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(int, int, string, string) []users.Domain); ok {
		r0 = rf(page, limit, sort, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, string, string) int); ok {
		r1 = rf(page, limit, sort, order)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(page, limit, sort, order)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: domain
func (_m *UseCase) Login(domain *users.Domain) (users.Domain, string, error) {
	ret := _m.Called(domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*users.Domain) string); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*users.Domain) error); ok {
		r2 = rf(domain)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: domain
func (_m *UseCase) Register(domain *users.Domain) (users.Domain, string, error) {
	ret := _m.Called(domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*users.Domain) string); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*users.Domain) error); ok {
		r2 = rf(domain)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: id, domain
func (_m *UseCase) Update(id primitive.ObjectID, domain *users.Domain) (users.Domain, error) {
	ret := _m.Called(id, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *users.Domain) users.Domain); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, *users.Domain) error); ok {
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

// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dtothreads "charum/dto/threads"

	mock "github.com/stretchr/testify/mock"

	pagination "charum/dto/pagination"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	threads "charum/business/threads"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// AdminDelete provides a mock function with given fields: threadID
func (_m *UseCase) AdminDelete(threadID primitive.ObjectID) (threads.Domain, error) {
	ret := _m.Called(threadID)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) threads.Domain); ok {
		r0 = rf(threadID)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(threadID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AdminUpdate provides a mock function with given fields: domain
func (_m *UseCase) AdminUpdate(domain *threads.Domain) (threads.Domain, error) {
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

// DeleteByThreadID provides a mock function with given fields: threadID
func (_m *UseCase) DeleteByThreadID(threadID primitive.ObjectID) error {
	ret := _m.Called(threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DomainToResponse provides a mock function with given fields: domain
func (_m *UseCase) DomainToResponse(domain threads.Domain) (dtothreads.Response, error) {
	ret := _m.Called(domain)

	var r0 dtothreads.Response
	if rf, ok := ret.Get(0).(func(threads.Domain) dtothreads.Response); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(dtothreads.Response)
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
func (_m *UseCase) DomainsToResponseArray(domains []threads.Domain) ([]dtothreads.Response, error) {
	ret := _m.Called(domains)

	var r0 []dtothreads.Response
	if rf, ok := ret.Get(0).(func([]threads.Domain) []dtothreads.Response); ok {
		r0 = rf(domains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtothreads.Response)
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

// GetAllByTopicID provides a mock function with given fields: topicID
func (_m *UseCase) GetAllByTopicID(topicID primitive.ObjectID) ([]threads.Domain, error) {
	ret := _m.Called(topicID)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) []threads.Domain); ok {
		r0 = rf(topicID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(topicID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllByUserID provides a mock function with given fields: userID
func (_m *UseCase) GetAllByUserID(userID primitive.ObjectID) ([]threads.Domain, error) {
	ret := _m.Called(userID)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) []threads.Domain); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
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

// GetLikedByUserID provides a mock function with given fields: userID
func (_m *UseCase) GetLikedByUserID(userID primitive.ObjectID) ([]threads.Domain, error) {
	ret := _m.Called(userID)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) []threads.Domain); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
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

// GetManyWithPagination provides a mock function with given fields: _a0, domain
func (_m *UseCase) GetManyWithPagination(_a0 pagination.Request, domain *threads.Domain) ([]threads.Domain, int, int, error) {
	ret := _m.Called(_a0, domain)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(pagination.Request, *threads.Domain) []threads.Domain); ok {
		r0 = rf(_a0, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(pagination.Request, *threads.Domain) int); ok {
		r1 = rf(_a0, domain)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(pagination.Request, *threads.Domain) int); ok {
		r2 = rf(_a0, domain)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(pagination.Request, *threads.Domain) error); ok {
		r3 = rf(_a0, domain)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Like provides a mock function with given fields: userID, threadID
func (_m *UseCase) Like(userID primitive.ObjectID, threadID primitive.ObjectID) error {
	ret := _m.Called(userID, threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// Unlike provides a mock function with given fields: userID, threadID
func (_m *UseCase) Unlike(userID primitive.ObjectID, threadID primitive.ObjectID) error {
	ret := _m.Called(userID, threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserUpdate provides a mock function with given fields: domain
func (_m *UseCase) UserUpdate(domain *threads.Domain) (threads.Domain, error) {
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

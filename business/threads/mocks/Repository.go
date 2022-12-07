// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	query "charum/dto/query"

	threads "charum/business/threads"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AppendLike provides a mock function with given fields: userID, threadID
func (_m *Repository) AppendLike(userID primitive.ObjectID, threadID primitive.ObjectID) error {
	ret := _m.Called(userID, threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckLikedByUserID provides a mock function with given fields: userID, threadID
func (_m *Repository) CheckLikedByUserID(userID primitive.ObjectID, threadID primitive.ObjectID) error {
	ret := _m.Called(userID, threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: domain
func (_m *Repository) Create(domain *threads.Domain) (threads.Domain, error) {
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

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id primitive.ObjectID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByUserID provides a mock function with given fields: id
func (_m *Repository) DeleteAllByUserID(id primitive.ObjectID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllByTopicID provides a mock function with given fields: topicID
func (_m *Repository) GetAllByTopicID(topicID primitive.ObjectID) ([]threads.Domain, error) {
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
func (_m *Repository) GetAllByUserID(userID primitive.ObjectID) ([]threads.Domain, error) {
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
func (_m *Repository) GetByID(id primitive.ObjectID) (threads.Domain, error) {
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
func (_m *Repository) GetLikedByUserID(userID primitive.ObjectID) ([]threads.Domain, error) {
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
func (_m *Repository) GetManyWithPagination(_a0 query.Request, domain *threads.Domain) ([]threads.Domain, int, error) {
	ret := _m.Called(_a0, domain)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(query.Request, *threads.Domain) []threads.Domain); ok {
		r0 = rf(_a0, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(query.Request, *threads.Domain) int); ok {
		r1 = rf(_a0, domain)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(query.Request, *threads.Domain) error); ok {
		r2 = rf(_a0, domain)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveLike provides a mock function with given fields: userID, threadID
func (_m *Repository) RemoveLike(userID primitive.ObjectID, threadID primitive.ObjectID) error {
	ret := _m.Called(userID, threadID)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(userID, threadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SuspendByUserID provides a mock function with given fields: domain
func (_m *Repository) SuspendByUserID(domain *threads.Domain) error {
	ret := _m.Called(domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*threads.Domain) error); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: domain
func (_m *Repository) Update(domain *threads.Domain) (threads.Domain, error) {
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

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

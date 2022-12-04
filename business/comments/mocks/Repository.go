// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	comments "charum/business/comments"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountByThreadID provides a mock function with given fields: threadID
func (_m *Repository) CountByThreadID(threadID primitive.ObjectID) (int, error) {
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
func (_m *Repository) Create(domain *comments.Domain) (comments.Domain, error) {
	ret := _m.Called(domain)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(*comments.Domain) comments.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*comments.Domain) error); ok {
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

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id primitive.ObjectID) (comments.Domain, error) {
	ret := _m.Called(id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) comments.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByThreadID provides a mock function with given fields: threadID
func (_m *Repository) GetByThreadID(threadID primitive.ObjectID) ([]comments.Domain, error) {
	ret := _m.Called(threadID)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) []comments.Domain); ok {
		r0 = rf(threadID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(threadID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: domain
func (_m *Repository) Update(domain *comments.Domain) (comments.Domain, error) {
	ret := _m.Called(domain)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(*comments.Domain) comments.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*comments.Domain) error); ok {
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

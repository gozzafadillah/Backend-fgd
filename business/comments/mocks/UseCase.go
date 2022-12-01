// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	comments "charum/business/comments"
	dto "charum/dto"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: domain
func (_m *UseCase) Create(domain *comments.Domain) (comments.Domain, error) {
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

// Delete provides a mock function with given fields: id, userID
func (_m *UseCase) Delete(id primitive.ObjectID, userID primitive.ObjectID) (comments.Domain, error) {
	ret := _m.Called(id, userID)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, primitive.ObjectID) comments.Domain); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, primitive.ObjectID) error); ok {
		r1 = rf(id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DomainToResponse provides a mock function with given fields: comment
func (_m *UseCase) DomainToResponse(comment comments.Domain) (dto.ResponseComment, error) {
	ret := _m.Called(comment)

	var r0 dto.ResponseComment
	if rf, ok := ret.Get(0).(func(comments.Domain) dto.ResponseComment); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Get(0).(dto.ResponseComment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comments.Domain) error); ok {
		r1 = rf(comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DomainToResponseArray provides a mock function with given fields: _a0
func (_m *UseCase) DomainToResponseArray(_a0 []comments.Domain) ([]dto.ResponseComment, error) {
	ret := _m.Called(_a0)

	var r0 []dto.ResponseComment
	if rf, ok := ret.Get(0).(func([]comments.Domain) []dto.ResponseComment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ResponseComment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]comments.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByThreadID provides a mock function with given fields: threadID
func (_m *UseCase) GetByThreadID(threadID primitive.ObjectID) ([]comments.Domain, error) {
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
func (_m *UseCase) Update(domain *comments.Domain) (comments.Domain, error) {
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

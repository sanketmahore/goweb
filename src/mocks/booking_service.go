// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "crud/src/domain"

	mock "github.com/stretchr/testify/mock"
)

// BookingService is an autogenerated mock type for the BookingService type
type BookingService struct {
	mock.Mock
}

// Create provides a mock function with given fields: bookingReq
func (_m *BookingService) Create(bookingReq *domain.Booking) error {
	ret := _m.Called(bookingReq)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Booking) error); ok {
		r0 = rf(bookingReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *BookingService) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *BookingService) Get(id string) (*domain.Booking, error) {
	ret := _m.Called(id)

	var r0 *domain.Booking
	if rf, ok := ret.Get(0).(func(string) *domain.Booking); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *BookingService) List() []*domain.Booking {
	ret := _m.Called()

	var r0 []*domain.Booking
	if rf, ok := ret.Get(0).(func() []*domain.Booking); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Booking)
		}
	}

	return r0
}

// Login provides a mock function with given fields: creds
func (_m *BookingService) Login(creds *domain.Credentials) (bool, error) {
	ret := _m.Called(creds)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.Credentials) bool); ok {
		r0 = rf(creds)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Credentials) error); ok {
		r1 = rf(creds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: booking
func (_m *BookingService) Update(booking *domain.Booking) error {
	ret := _m.Called(booking)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Booking) error); ok {
		r0 = rf(booking)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

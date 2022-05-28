package service

import (
	"crud/src/domain"
	"crud/src/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListBooking(t *testing.T) {
	dao := &mocks.BookingDao{}
	service := NewBookingService(dao)
	cases := []struct {
		name string
		res  []*domain.Booking
	}{

		{
			"success",
			[]*domain.Booking{},
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				dao.On("ReturnAllBookings").Return(c.res)
				res := service.List()
				assert.Equal(t, c.res, res)
			},
		)
	}

}

func TestCreateBooking(t *testing.T) {
	dao := &mocks.BookingDao{}
	service := NewBookingService(dao)
	cases := []struct {
		name       string
		bookingReq *domain.Booking
		err        error
	}{

		{
			"success",
			&domain.Booking{Id: 1, User: "mock", Members: 3},
			nil,
		},
		{
			"duplicate",
			&domain.Booking{Id: 1, User: "mock", Members: 3},
			domain.ErrConflict,
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				dao.On("CreateBooking", mock.Anything).Return(c.err).Once()
				
				err := service.Create(c.bookingReq)
				assert.Equal(t, c.err, err)
			},
		)
	}
}

package bookingimpl

import (
	"crud/src/domain"
	"crud/src/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
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

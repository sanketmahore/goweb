package bookingimpl

import (
	"crud/src/domain"
	"testing"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMysqlGet(t *testing.T) {
	dao := NewBookingDao()
	booking := []struct {
		name string
		id   string
		res  *domain.Booking
		err  error
	}{
		{
			"success",
			"1",
			&domain.Booking{Id: 2, User: "Mytest-1", Members: 5},
			nil,
		},
		{
			"failure",
			"100",
			nil,
			domain.ErrNotFound,
		},
	}
	for _, c := range booking {
		t.Run(c.name, func(t *testing.T) {
			r, err := dao.ReturnSingleBooking(c.id)
			assert.Equal(t, c.err, errors.Cause(err))
			if c.err == nil {
				assert.NotEmpty(t, r)
			}
		})
	}
}

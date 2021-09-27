package dao

import (
	"crud/src/domain"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMysqlSave(t *testing.T) {
	dao := NewBookingDao()
	booking := []struct {
		name    string
		booking *domain.Booking
		err     error
	}{
		{
			"success",
			&domain.Booking{Id: 100, User: "Mytest", Members: 5},
			nil,
		},
		{
			"save-conflict",
			&domain.Booking{Id: 100, User: "Mytest", Members: 5},
			domain.ErrConflict,
		},
	}
	for _, c := range booking {
		t.Run(c.name, func(t *testing.T) {
			err := dao.CreateBooking(c.booking)
			assert.Equal(t, c.err, errors.Cause(err))			
		})
	}
}
func TestMysqlGetSingleRecord(t *testing.T) {
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
			&domain.Booking{Id: 1, User: "Mytest", Members: 5},
			nil,
		},
		{
			"failure",
			"1000",
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

func TestMysqlGetAll(t *testing.T) {
	dao := NewBookingDao()
	booking := []struct {
		name string
		res  *domain.Booking
		err  error
	}{
		{
			"success",
			&domain.Booking{},
			nil,
		},
	}
	for _, c := range booking {
		t.Run(c.name, func(t *testing.T) {
			r := dao.ReturnAllBookings()
			assert.NotEmpty(t, r)
		})
	}
}

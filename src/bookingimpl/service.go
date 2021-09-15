package bookingimpl

import (
	"crud/src/domain"
	"strconv"	
)

func Create(bookingReq *domain.Booking) error{
	err := CreateBooking(bookingReq)
	return err
}

func List() []*domain.Booking{
	bookings := ReturnAllBookings()
	return bookings
}

func Get(id string) (*domain.Booking, error){
	booking, err := ReturnSingleBooking(id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func Update(booking *domain.Booking) error{
	_, err := Get(strconv.Itoa(booking.Id))
	if err != nil {
		return err
	}
	err = UpdateBooking(booking)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id string) error{
	err := DeleteBooking(id)
	if err != nil {
		return err
	}
	return nil
}
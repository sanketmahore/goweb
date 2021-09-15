package bookingimpl

import (
	"crud/src/domain"
	"strconv"
)

var bookingDao domain.BookingDao

type service struct{}

func NewBookingService(dao domain.BookingDao) domain.BookingService {
	bookingDao = dao
	return &service{}
}

func (*service) Create(bookingReq *domain.Booking) error {
	err := bookingDao.CreateBooking(bookingReq)
	return err
}

func (*service) List() []*domain.Booking {
	bookings := bookingDao.ReturnAllBookings()
	return bookings
}

func (*service) Get(id string) (*domain.Booking, error) {
	booking, err := bookingDao.ReturnSingleBooking(id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *service) Update(booking *domain.Booking) error {
	_, err := s.Get(strconv.Itoa(booking.Id))
	if err != nil {
		return err
	}
	err = bookingDao.UpdateBooking(booking)
	if err != nil {
		return err
	}
	return nil
}

func (*service) Delete(id string) error {
	err := bookingDao.DeleteBooking(id)
	if err != nil {
		return err
	}
	return nil
}

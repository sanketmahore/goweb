package service

import (
	"goweb_microservices/admin_microservice/src/domain"
	"strconv"

)

//var bookingDao domain.BookingDao

type service struct {
	bookingDao domain.BookingDao
}

func NewBookingService(dao domain.BookingDao) domain.BookingService {
	return &service{bookingDao: dao}
}

func (s *service) Create(bookingReq *domain.Booking) error {
	err := s.bookingDao.CreateBooking(bookingReq)
	return err
}

func (s *service) List() []*domain.Booking {
	bookings := s.bookingDao.ReturnAllBookings()
	return bookings
}

func (s *service) Get(id string) (*domain.Booking, error) {
	booking, err := s.bookingDao.ReturnSingleBooking(id)
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
	err = s.bookingDao.UpdateBooking(booking)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id string) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	err = s.bookingDao.DeleteBooking(id)
	if err != nil {
		return err
	}
	return nil
}

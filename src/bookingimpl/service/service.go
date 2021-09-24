package service

import (
	"context"
	"crud/src/authentication"
	"crud/src/domain"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

//var bookingDao domain.BookingDao

type service struct {
	bookingDao domain.BookingDao
}

func NewBookingService(dao domain.BookingDao) domain.BookingService {
	return &service{bookingDao: dao}
}

func (s *service) Login(creds *domain.Credentials) (bool, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return false, fmt.Errorf("failed to connect %v", err)
	}

	defer conn.Close()

	c := authentication.NewLoginServiceClient(conn)
	req := &authentication.Credentials{
		Username: creds.Username,
		Password: creds.Password,
	}

	res, err := c.Authenticate(context.Background(), req)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, fmt.Errorf("error in RPC %v", err)
	}

	return res.Outcome, nil
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

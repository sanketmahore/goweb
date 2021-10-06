package domain

import (
	"goweb_microservices/login_microservice/authentication"
	"net/http"
)

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Controller interface {
	HomePage(w http.ResponseWriter, r *http.Request)
	CreateNewBooking(w http.ResponseWriter, r *http.Request)
	GetAllBookings(w http.ResponseWriter, r *http.Request)
	GetSingleBooking(w http.ResponseWriter, r *http.Request)
	UpdateBooking(w http.ResponseWriter, r *http.Request)
	DeleteBooking(w http.ResponseWriter, r *http.Request)
}

type BookingService interface {
	Create(bookingReq *Booking) error
	Update(booking *Booking) error
	Get(id string) (*Booking, error)
	List() []*Booking
	Delete(id string) error
}

type BookingDao interface {
	CreateBooking(booking *Booking) error
	ReturnAllBookings() []*Booking
	ReturnSingleBooking(id string) (*Booking, error)
	UpdateBooking(booking *Booking) error
	DeleteBooking(id string) error
	AuthenticateUser(creds *authentication.Credentials) (bool, error)
}

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrConflict      = Error("record already exists")
	ErrNotFound      = Error("not found")
	ErrInvalidSyntax = Error("invalid syntax")
)

const (
	MysqlDupicate = 1062
)

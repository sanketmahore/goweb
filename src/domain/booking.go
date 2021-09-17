package domain

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
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
}

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrConflict = Error("record already exists")
	ErrNotFound = Error("not found")
)

const(
	MysqlDupicate = 1062
)
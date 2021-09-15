package domain

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

type Error string

func (e Error) Error() string{
	return string(e)
}

var (
	ErrConflict = Error("Duplicate entry")
	ErrNotFound = Error("not found")
)

// type BookingService interface {
// 	Create(bookingReq *Booking) Error
// 	Update(bookingReq *Booking)
// 	Get(id string)
// 	GetAll()
// }

// type BookingDao interface {
// 	Create(bookingReq *Booking) Error
// 	Update(bookingReq *Booking)
// 	Get(id string)
// 	GetAll()
// }

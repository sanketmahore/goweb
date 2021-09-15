package bookingimpl

import (
	"crud/src/domain"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var err error

type dao struct{}

func NewBookingDao() domain.BookingDao {
	return &dao{}
}
func init() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/booking_api?charset=utf8&parseTime=True")
	db.AutoMigrate(&domain.Booking{})
}

func (*dao) CreateBooking(booking *domain.Booking) error {
	err := db.Create(booking).Error
	return err
}

func (*dao) ReturnAllBookings() []*domain.Booking {
	bookings := make([]*domain.Booking, 0)
	db.Find(&bookings)
	return bookings
}

func (*dao) ReturnSingleBooking(id string) (*domain.Booking, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	booking := &domain.Booking{
		Id: ID,
	}
	err = db.First(booking).Error
	if gorm.IsRecordNotFoundError(err) {
		err = domain.ErrNotFound
		return nil, err
	}
	return booking, nil
}

func (*dao) UpdateBooking(booking *domain.Booking) error {
	//db.Exec("UPDATE Bookings SET User =? WHERE Id=?", booking.User, booking.Id)
	err := db.Model(booking).Update(booking).Error
	if err != nil {
		fmt.Println("Mysql Error..", err)
		return err
	}
	return nil
}

func (*dao) DeleteBooking(id string) error {
	//db.Exec("DELETE FROM Bookings WHERE Id=?", id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	r := db.Delete(domain.Booking{Id: ID})
	if r.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

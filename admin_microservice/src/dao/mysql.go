package dao

import (
	"goweb_microservices/admin_microservice/src/domain"
	"goweb_microservices/login_microservice/authentication"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/go-sql-driver/mysql"
)

type dao struct {
	session *gorm.DB
}

// func NewBookingDao() domain.BookingDao {
// 	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/booking_api?charset=utf8&parseTime=True")
// 	if err != nil {
// 		println("Mysql connection failed..")
// 	}
// 	db.AutoMigrate(&domain.Booking{})
// 	db.AutoMigrate(&domain.Credentials{})
// 	return &dao{session: db}
// }

func NewBookingDao() domain.BookingDao {
	db, err := ConnectLoop("mysql",
		"root:root@tcp(172.16.238.10:3306)/booking_api?charset=utf8&parseTime=True",
		30*time.Second)
	if err != nil {
		println("Mysql connection failed.. ", err.Error())
	} else {
		println("Mysql connection Successful")
	}
	db.AutoMigrate(&domain.Booking{})
	db.AutoMigrate(&domain.Credentials{})
	return &dao{session: db}
}

func (d *dao) CreateBooking(booking *domain.Booking) error {
	err := d.session.Create(booking).Error
	switch err := err.(type) {
	case *mysql.MySQLError:
		if err.Number == domain.MysqlDupicate {
			return domain.ErrConflict
		}
		return err
	}
	return err
}

func (d *dao) ReturnAllBookings() []*domain.Booking {
	bookings := make([]*domain.Booking, 0)
	d.session.Find(&bookings)
	return bookings
}

func (d *dao) ReturnSingleBooking(id string) (*domain.Booking, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, domain.ErrInvalidSyntax
	}
	booking := &domain.Booking{
		Id: ID,
	}
	err = d.session.First(booking).Error
	if gorm.IsRecordNotFoundError(err) {
		err = domain.ErrNotFound
		return nil, err
	}
	return booking, nil
}

func (d *dao) UpdateBooking(booking *domain.Booking) error {
	//db.Exec("UPDATE Bookings SET User =? WHERE Id=?", booking.User, booking.Id)
	err := d.session.Model(booking).Update(booking).Error
	if err != nil {
		fmt.Println("Mysql Error..", err)
		return err
	}
	return nil
}

func (d *dao) DeleteBooking(id string) error {
	//db.Exec("DELETE FROM Bookings WHERE Id=?", id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	r := d.session.Delete(domain.Booking{Id: ID})
	if r.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (d *dao) AuthenticateUser(creds *authentication.Credentials) (bool, error) {
	fmt.Println("In mysql : AuthenticateUser start")
	credentials := &domain.Credentials{
		Username: creds.Username,
		Password: creds.Password,
	}

	//err := d.session.First(credentials).Error
	err := d.session.First(credentials, "username = ? and password = ?", creds.Username, creds.Password).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = domain.ErrNotFound
			return false, err
		}
		return false, err
	}

	fmt.Println("In mysql : AuthenticateUser end")
	return true, nil
}

func ConnectLoop(driver, DSN string, timeout time.Duration) (*gorm.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)
		case <-ticker.C:
			db, err := gorm.Open(driver, DSN)
			if err == nil {
				fmt.Println("!!! DB Connection SUCCESSFUL !!!")
				return db, nil
			}
			fmt.Println(errors.Wrapf(err, "failed to connect to db %s", DSN))
		}
	}
}

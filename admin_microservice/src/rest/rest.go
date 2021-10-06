package rest

import (
	"goweb_microservices/admin_microservice/src/dao"
	"goweb_microservices/admin_microservice/src/service"
	"goweb_microservices/admin_microservice/src/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type bookingController struct {
	service domain.BookingService
}

func NewBookingController(s domain.BookingService) *bookingController {
	return &bookingController{service: s}
}

func HandleRequests() {

	log.Println("Starting development server at http://127.0.0.1:10000/")

	log.Println("Quit the server with CONTROL-C.")

	bookingController := NewBookingController(service.NewBookingService(dao.NewBookingDao()))

	// creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", bookingController.HomePage)
	myRouter.HandleFunc("/new-booking", bookingController.CreateNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", bookingController.GetAllBookings).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", bookingController.GetSingleBooking).Methods("GET")
	myRouter.HandleFunc("/update", bookingController.UpdateBooking).Methods("PATCH")
	myRouter.HandleFunc("/delete/{id}", bookingController.DeleteBooking).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func (c *bookingController) HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to HomePage!!!!!!!!!!!!")

	fmt.Println("Endpoint Hit: HomePage")
}

func (c *bookingController) CreateNewBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking *domain.Booking
	err := json.Unmarshal(reqBody, &booking)
	if err != nil {
		http.Error(w, "Can’t unmarshal JSON object into struct", http.StatusBadRequest)
		return
	}
	err = c.service.Create(booking)
	if err != nil {
		if errors.Cause(err) == domain.ErrConflict {
			http.Error(w, "booking already exists", http.StatusConflict)
			return
		}
		http.Error(w, "unable to create booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (c *bookingController) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := c.service.List()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (c *bookingController) GetSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	booking, err := c.service.Get(id)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		if errors.Cause(err) == domain.ErrInvalidSyntax {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}
		http.Error(w, "unable to get booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (c *bookingController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking *domain.Booking
	json.Unmarshal(reqBody, &booking)
	err := c.service.Update(booking)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid syntax") {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}
		http.Error(w, "unable to update booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (c *bookingController) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.service.Delete(id)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid syntax") {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}
		http.Error(w, "unable to delete booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	http.Error(w, "booking deleted", http.StatusNoContent)
	json.NewEncoder(w).Encode("Record is deleted...")
}

package bookingimpl

import (
	"crud/src/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

//type bookingService domain.BookingService

var bookingService domain.BookingService

func HandleRequests() {

	log.Println("Starting development server at http://127.0.0.1:10000/")

	log.Println("Quit the server with CONTROL-C.")

	bookingService = NewBookingService(NewBookingDao())

	// creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", returnAllBookings).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")
	myRouter.HandleFunc("/update", updatedBooking).Methods("PATCH")
	myRouter.HandleFunc("/delete/{id}", deleteBooking).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to HomePage!!!!!!!!!!!!")

	fmt.Println("Endpoint Hit: HomePage")
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking *domain.Booking
	err := json.Unmarshal(reqBody, &booking)
	if err != nil {
		http.Error(w, "Can’t unmarshal JSON object into struct", http.StatusBadRequest)
		return
	}
	err = bookingService.Create(booking)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
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

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := bookingService.List()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	booking, err := bookingService.Get(id)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid syntax") {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}
		http.Error(w, "unable to get booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func updatedBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking *domain.Booking
	json.Unmarshal(reqBody, &booking)
	err := bookingService.Update(booking)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		http.Error(w, "unable to update booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := bookingService.Delete(id)
	if err != nil {
		if errors.Cause(err) == domain.ErrNotFound {
			http.Error(w, "booking not found with id", http.StatusNotFound)
			return
		}
		http.Error(w, "unable to delete booking", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	http.Error(w, "booking deleted", http.StatusNoContent)
	json.NewEncoder(w).Encode("Record is deleted...")
}

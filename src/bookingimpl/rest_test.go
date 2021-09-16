package bookingimpl

import (
	"crud/src/mocks"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.ServeHTTP(rr, req)
	return rr
}

func TestGetAllBookings(t *testing.T) {
	service := &mocks.BookingService{}
	//controller := NewBookingController(service)
	bookings := []struct {
		name   string
		res    string
		status int
	}{
		{
			"success",
			"",
			200,
		},
	}
	for _, c := range bookings {
		t.Run(c.name, func(t *testing.T) {
			service.On("List").Return(c.res).Once()

			httpRequest, err := http.NewRequest("GET", "http://127.0.0.1:10000/all-bookings", nil)
			if err != nil {
				log.Println("Unable to create GET request", err.Error())
			}

			httpRequest.Header.Set("Content-Type", "application/json")
			httpRequest.Header.Set("Accept", "application/json")

			httpWriter := executeRequest(httpRequest)

			res := httpWriter.Result()
			body, _ := ioutil.ReadAll(res.Body)
			assert.Equal(t, c.status, httpWriter.Code)
			assert.Equal(t, c.res, string(body))
		},
		)
	}

}

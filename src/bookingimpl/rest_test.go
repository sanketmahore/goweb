package bookingimpl

import (
	"crud/src/domain"
	"crud/src/mocks"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func executeRequest(restFunction func(http.ResponseWriter, *http.Request), httpRequest *http.Request) *httptest.ResponseRecorder {
	httpWriter := httptest.NewRecorder()
	handler := http.HandlerFunc(restFunction)
	handler.ServeHTTP(httpWriter, httpRequest)
	return httpWriter
}

func TestGetAllBookings(t *testing.T) {
	service := &mocks.BookingService{}
	controller := NewBookingController(service)
	bookings := []struct {
		name        string
		service_res []*domain.Booking
		res         string
		status      int
	}{
		{
			"success",
			[]*domain.Booking{
				{
					Id: 1, User: "mock", Members: 3,
				},
				{
					Id: 2, User: "mock", Members: 3,
				},
				{
					Id: 3, User: "mock", Members: 3,
				},
			},
			"[{\"id\":1,\"user\":\"mock\",\"members\":3},{\"id\":2,\"user\":\"mock\",\"members\":3},{\"id\":3,\"user\":\"mock\",\"members\":3}]\n",
			200,
		},
	}
	for _, c := range bookings {
		t.Run(c.name, func(t *testing.T) {
			service.On("List").Return(c.service_res).Once()

			httpRequest, err := http.NewRequest("GET", "http://127.0.0.1:10000/all-bookings", nil)
			if err != nil {
				log.Println("Unable to create GET request", err.Error())
			}

			httpRequest.Header.Set("Content-Type", "application/json")
			httpRequest.Header.Set("Accept", "application/json")

			httpWriter := executeRequest(controller.GetAllBookings, httpRequest)
			res := httpWriter.Result()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Panic("Unable to read response " + err.Error())
			}
			assert.Equal(t, c.status, httpWriter.Code)
			assert.Equal(t, c.res, string(body))
		},
		)
	}

}

func TestGetSingleBooking(t *testing.T) {
	service := &mocks.BookingService{}
	controller := NewBookingController(service)

	cases := []struct {
		name     string
		id       string
		mock_res *domain.Booking
		mock_err error
		res      string
		status   int
	}{
		{
			"success: record found",
			"1",
			&domain.Booking{Id: 1, User: "mock", Members: 3},
			nil,
			"{\"id\":1,\"user\":\"mock\",\"members\":3}\n",
			http.StatusOK,
		},
		{
			"fail: record not found",
			"2",
			nil,
			domain.ErrNotFound,
			"booking not found with id\n",
			http.StatusNotFound,
		},
		{
			"fail: invalid id",
			"a",
			nil,
			domain.ErrInvalidSyntax,
			"incorrect id\n",
			http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			service.On("Get", mock.Anything).Return(c.mock_res, c.mock_err).Once()
			httpRequest, err := http.NewRequest("GET", "http://127.0.0.1:10000/booking", nil)
			if err != nil {
				log.Println("Unable to create GET request", err.Error())
			}
			q := httpRequest.URL.Query()
			q.Add("id", c.id)
			httpRequest.URL.RawQuery = q.Encode()
			httpRequest.Header.Set("Content-Type", "application/json")
			httpRequest.Header.Set("Accept", "application/json")
			httpWriter := executeRequest(controller.GetSingleBooking, httpRequest)
			res := httpWriter.Result()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Panic("Unable to read response " + err.Error())
			}
			assert.Equal(t, c.status, httpWriter.Code)
			assert.Equal(t, c.res, string(body))
		},
		)
	}
}

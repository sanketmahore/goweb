package rest

import (
	"goweb_microservices/login_microservice/domain"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"goweb_microservices/login_microservice/service"

	"github.com/gorilla/mux"
)

type loginController struct {
	service domain.LoginService
}

func NewLoginController(s domain.LoginService) *loginController {
	return &loginController{service: s}
}

func HandleRequests() {

	log.Println("Starting development server at http://127.0.0.1:10001/")

	log.Println("Quit the server with CONTROL-C.")

	loginController := NewLoginController(service.NewLoginService())

	// creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/login", loginController.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func (c *loginController) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var creds *domain.Credentials
	err := json.Unmarshal(reqBody, &creds)
	if err != nil {
		http.Error(w, "Can’t unmarshal JSON object into struct", http.StatusBadRequest)
		return
	}
	auth_res, err := c.service.Login(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	if !auth_res {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(auth_res)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth_res)
}

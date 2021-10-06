package domain

import (
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
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginService interface {
	Login(creds *Credentials) (bool, error)
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

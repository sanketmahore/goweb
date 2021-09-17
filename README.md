# goweb
* Build:
go build
run exe

* Importent packages:
"github.com/gorilla/mux"
"net/http"
"github.com/jinzhu/gorm"
"github.com/go-sql-driver/mysql"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"


* Unit tests:
install
    go get github.com/vektra/mockery/v2/.../
Generate mock interface:
    mockery --case snake --name BookingService --dir domain

* Gorm Docs:
https://gorm.io/docs/

* Mod commands:
go mod init learn
go mod tidy
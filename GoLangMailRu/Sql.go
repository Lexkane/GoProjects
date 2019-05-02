package main

import (
	"fmt"
	"net/http"
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func main(){



	var err error


	db,err=sql.Open("mysql","root@tcp(localhost:3306)/msu-go-11?charset=utf8")
	PanicOnErr(err)
}

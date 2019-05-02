package main

import (
	"fmt"
	"net/http"
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func main(){
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err = http.ErrNoCookie {
			w.Write([] byte (loginFormTmpl))
			return
		} else if err != nil {

			PanicOnErr(err)
		}
		fmt.Fpring(w, "Welcome, "+sessionID.Value)
	})
	}

http.HandleFunc("/get_cookie", func (w http.ResponseWriter, r *http.Request)) {
	r.Parsefrom()
	inputLogin := r.Form["login"][0]
	expiration := time.Now().Add(365 * 24 * time.Hour)

	sessionId:=RandStringRunes(32)
	session[sessionID]=inputLogin

	cookie := http.Cookie{
		Name:    "seesion_id",
		Value:   inputLogin,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
})



}
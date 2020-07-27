package main

import (
    "database/sql"
  //  "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    // 大文字だと Public 扱い
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type JWT struct {
    Token string `json:"token"`
}

type Error struct {
    Message string `json:"message"`
}

func signup(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("successfully called signup"))
}

func login(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("successfully called login"))
    db, err := sql.Open("mysql", "root:hoge@tcp(172.27.145.106:3306)/users")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    temp_email := "mail"

    var email string
    var password string

    err = db.QueryRow("SELECT email FROM users WHERE email = ?", temp_email).Scan(&email, &password)
    if err != nil {
        log.Fatal(err)
    }
   // w.Write([]byte(email))
   // w.Write([]byte(password))
}

var db *sql.DB

func main() {
    db, err := sql.Open("mysql", "root:hoge@/users")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	// urls.py
	router := mux.NewRouter()

    // endpoint
    router.HandleFunc("/signup", signup).Methods("POST")
    router.HandleFunc("/login", login).Methods("POST")
	
	// console に出力する
	log.Println("サーバー起動 : 8000 port で受信")

	// log.Fatal は、異常を検知すると処理の実行を止めてくれる
	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/davecgh/go-spew/spew"
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

func errorInResponse(w http.ResponseWriter, status int, error Error) {
    w.WriteHeader(status) // 400 とか 500 などの HTTP status コードが入る
    json.NewEncoder(w).Encode(error)
    return
}

func signup(w http.ResponseWriter, r *http.Request) {
    var user User
    var error Error

    // r.body に何が帰ってくるか確認
    fmt.Println(r.Body)

    // https://golang.org/pkg/encoding/json/#NewDecoder
    json.NewDecoder(r.Body).Decode(&user)

    if user.Email == "" {
        error.Message = "Email は必須です。"
        errorInResponse(w, http.StatusBadRequest, error)
        return
    }

    if user.Password == "" {
        error.Message = "パスワードは必須です。"
        errorInResponse(w, http.StatusBadRequest, error)
        return
    }

    // user に何が格納されているのか
    fmt.Println(user)

    // dump も出せる
    fmt.Println("---------------------")
    spew.Dump(user)

}

func login(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("successfully called login"))

}

var db *sql.DB

func main() {
    db, err := sql.Open("mysql", "root:hoge@/authdb")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // temp_email := "mail"

    // var email string
    // var password string

    // err = db.QueryRow("SELECT email, passeord FROM users WHERE email = ?", temp_email).Scan(&email, &password)
    // if err != nil {
    //     log.Fatal(err)
    // }

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

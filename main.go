package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "golang.org/x/crypto/bcrypt"
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

func responseByJSON(w http.ResponseWriter, data interface{}) {
    json.NewEncoder(w).Encode(data)
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
        error.Message = "Email は必須です．"
        errorInResponse(w, http.StatusBadRequest, error)
        return
    }

    if user.Password == "" {
        error.Message = "パスワードは必須です．"
        errorInResponse(w, http.StatusBadRequest, error)
        return
    }


    fmt.Println("---------------------")
    
    // パスワードのハッシュを生成
    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("パスワード: ", user.Password)
    fmt.Println("ハッシュ化されたパスワード", hash)

    user.Password = string(hash)
    fmt.Println("コンバート後のパスワード: ", user.Password)


    // データベースに接続
    db, err := sql.Open("mysql", "root:hoge@/authdb")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    // query 発行
    // Scan で、Query 結果を変数に格納
    ins, err := db.Prepare("INSERT INTO users(email, password) VALUES(?, ?);")
    if err != nil {
        error.Message = "データベース処理に失敗しました"
        errorInResponse(w, http.StatusInternalServerError, error)
        return
    }
    
    ins.Exec(user.Email, user.Password)

    defer ins.Close()

    // DB に登録できたらパスワードをからにしておく
    user.Password = ""
    w.Header().Set("Content-Type", "application/json")

    // JSON 形式で結果を返却
    responseByJSON(w, user)

    defer db.Close()
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

    defer db.Close()

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

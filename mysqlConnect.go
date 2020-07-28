package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)


type Error struct {
    Message string `json:"message"`
}


func main() {
    var error Error

    db, err := sql.Open("mysql", "root:hoge@/authdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ins, err := db.Prepare("INSERT INTO users(email, password) VALUES(?, ?);")
    if err != nil {
        error.Message = "データベース処理に失敗しました"
        return
    }
    ins.Exec("hoge", "aaaaaaa")

    // id := 4
    var id int
    var email string
    var password string
    err = db.QueryRow("SELECT id, email, password FROM users WHERE email = 'hoge'").Scan(&id, &email, &password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(id,email, password)


}
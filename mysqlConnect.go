package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:hoge@/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    id := 2
    var email string
    err = db.QueryRow("SELECT id, email FROM test WHERE id = ?", id).Scan(&id, &email)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(id,email)
}
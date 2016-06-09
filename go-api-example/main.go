package main

import (
    "net/http"
    "database/sql"
    "time"
    _ "github.com/lib/pq"
    "fmt"
)

func response(rw http.ResponseWriter, request *http.Request) {
    fmt.Println("inserting ...")
    str := fmt.Sprintf("inserted the user: %d", insertUser(fmt.Sprintf("tmp: %d", time.Now().Nanosecond())))
    rw.Write([]byte(str))
}

func main() {
    http.HandleFunc("/insert", response)
    http.HandleFunc("/proc", func(rw http.ResponseWriter, request *http.Request){
        i := 10000;
        var x float64;
        for ; i > 0; i-- {
            x = 3.14 * float64(i);
        }
        x += 1;
    });
    http.ListenAndServe(":3000", nil)
}


func getConnection()(*sql.DB){
    db, err := sql.Open("postgres",  "postgres://root:root@postgresql-server.dev/go-api-example?sslmode=disable")
    checkErr(err)
    return db;
}

func insertUser(username string)(id int64){
    db := getConnection();
    defer db.Close()
    error := db.QueryRow("INSERT INTO users (username) VALUES ($1) RETURNING id", username).Scan(&id)
    checkErr(error)
    return id;
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
package main

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
    "github.com/thiagorobert/thiagogo_util"
    "github.com/garyburd/redigo/redis"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<HTML><H1><FONT COLOR='red'>Jennifer, I love you!!!!i</FONT></H1></HTML>")
}

func add_timestamp(w http.ResponseWriter, r *http.Request) {
    timestamp := "test"
    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        panic(err)
    }
    defer c.Close()

    //set
    c.Do("SET", timestamp, timestamp)

    //get
    word, err := redis.String(c.Do("GET", timestamp))
    if err != nil {
        fmt.Println("key not found")
    }

    io.WriteString(w, "<HTML><H1><FONT COLOR='red'>" + word + "</FONT></H1></HTML>")
}

func main() {
    port := 8042
    fmt.Printf("starting server at port " + strconv.Itoa(port) + "\n")
    fmt.Printf(thiagogo_util.Reverse("Thiago Robert"))
    http.HandleFunc("/", hello)
    http.HandleFunc("/add", add_timestamp)
    http.ListenAndServe(":" + strconv.Itoa(port), nil)
}


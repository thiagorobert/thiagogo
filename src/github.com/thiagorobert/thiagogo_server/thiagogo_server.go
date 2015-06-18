package main

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
    "github.com/thiagorobert/thiagogo_util"
//    "github.com/garyburd/redigo/redis"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
//    n, err := conn.Do("APPEND", "key", "value")
    port := 8042
    fmt.Printf("starting server at port " + strconv.Itoa(port) + "\n")
    fmt.Printf(thiagogo_util.Reverse("Thiago Robert"))
    http.HandleFunc("/", hello)
    // http.ListenAndServe(":" + strconv.Itoa(port), nil)
}


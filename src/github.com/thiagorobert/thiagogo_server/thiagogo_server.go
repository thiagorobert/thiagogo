package main

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
    "github.com/thiagorobert/thiagogo_util"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    port := 8042
    fmt.Printf("starting server at port " + strconv.Itoa(port) + "\n")
    fmt.Printf(thiagogo_util.Reverse("Thiago Robert"))
    http.HandleFunc("/", hello)
    http.ListenAndServe(":" + strconv.Itoa(port), nil)
}


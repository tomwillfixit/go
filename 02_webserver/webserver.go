//usr/bin/env go run $0 $@; exit //allows you to run ./fs.go

// Simple Go webserver Hello World example 

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    fmt.Println("Starting Simple File Server @ http://localhost:8888")
    http.Handle("/", http.FileServer(http.Dir("./")))
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
                log.Fatal(err)
        }
}

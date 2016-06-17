//usr/bin/env go run $0 $@; exit //allows you to run ./fs.go

// Simple Go webserver Hello World example 

package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello World from Simple Go Webserver")
}

func main() {
 
    fmt.Println("Starting Simple Go Webserver @ http://localhost:8888")

    // Write Hello World message
    http.HandleFunc("/", helloHandler)

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
                log.Fatal(err)
        }
}

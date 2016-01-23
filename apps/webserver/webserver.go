package main

import (
  "net/http"
  "fmt"
  "log"
)

func main() {
  http.HandleFunc("/", home)
  fmt.Println("server started")

  err := http.ListenAndServe("localhost:2000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Simple web server")
}

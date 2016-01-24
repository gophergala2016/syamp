package main

import (
  "net/http"
  "fmt"
  "log"
  "time"
  "math/rand"
  "io/ioutil"
)
// this is just a test server it prints to Stdout
func main() {
  http.HandleFunc("/", home)
  fmt.Println("server started")
  log_data := applog()
  go func() {
    for {
      err := ioutil.WriteFile("log/log.txt", <-log_data, 777)
      if err != nil {
        fmt.Println(err)
      }
    }
  }()

  err := http.ListenAndServe("localhost:2000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Simple web server")
}

// Generator returns a channel
func applog() <-chan string {
  c := make(chan string)
  go func() {
    for {
      c <- fmt.Sprintf("%s", "http://localhost:2000")
      time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
    }
  }()

  return c
}

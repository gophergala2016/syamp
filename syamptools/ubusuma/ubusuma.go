package ubusuma

import (
  "os/exec"
  "fmt"
  "time"
  "math/rand"
  "log"
  "bufio"
)


// Generator
func Webapp() <-chan string {
  c := make(chan string)
  go func() {
    comm := "./app/webserver/webserver"

    cmd := exec.Command(comm)
    byt, err := cmd.StdoutPipe()
    if err != nil {
      fmt.Println(err)
    }
    if err := cmd.Start(); err != nil {
      log.Fatal(err)
    }
    for {
      scanner := bufio.NewScanner(byt)
	    go func() {
        for scanner.Scan() {
          c <- scanner.Text()
        }
	    }()
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()

  return c
}

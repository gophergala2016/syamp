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
    comm := "./apps/webserver/webserver"

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

// Generator show running programs
func Runningapps() <-chan string {
  c := make(chan string)
  go func() {
    comm := "ps"
    args := []string {
      "-a",
    }

    cmd := exec.Command(comm, args...)
    byt, err := cmd.Output()
    if err != nil {
      fmt.Println(err)
    }
    c <- fmt.Sprintf("%s", byt)
  }()

  return c
}

// kill process
func Term(pid string) <-chan string {
  c := make(chan string)
  go func() {
    comm := "kill"
    var args []string
    args = append(args, "-KILL")
    args = append(args, pid)

    cmd := exec.Command(comm, args...)
    err := cmd.Run()
    if err != nil {
      fmt.Println(err)
    }
    c <- fmt.Sprintf("%s", "Killed")
  }()

  return c
}

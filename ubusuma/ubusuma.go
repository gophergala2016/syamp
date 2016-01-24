package ubusuma

import (
  "os/exec"
  "fmt"
  "bufio"
  "strings"
)

// returns a slice of strings containing the short Description about
// the base operationg system.
func Metal() <-chan [4]string {
  c := make(chan [4]string)
  go func() {
    lsb := "lsb_release"
    args := []string {
      "-a",
    }
    cmd := exec.Command(lsb, args ...)
    byt, err := cmd.Output()
    var empty [4]string
    if err != nil {
      c <- empty
    }
    file := strings.NewReader(fmt.Sprintf("%s", byt))
    var kam [4]string
    var count int = 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      kam[count] = scanner.Text()
      count += 1
    }
    if err := scanner.Err(); err != nil {
      c <- empty
    }
    c <- kam
  }()

  return c
}

// Generator show running programs
func RunningUser() <-chan string {
  c := make(chan string)
  go func() {
    comm := "ps"
    args := []string {
      "-a",
    }
    cmd := exec.Command(comm, args...)
    byt, err := cmd.Output()
    if err != nil {
      c <- fmt.Sprintf("%s", err)
    }
    c <- fmt.Sprintf("%s", byt)
  }()

  return c
}


// kill process
func Kill(pid string) <-chan string {
  c := make(chan string)
  go func() {
    comm := "kill"
    var args []string
    args = append(args, "-KILL")
    args = append(args, pid)

    cmd := exec.Command(comm, args...)
    err := cmd.Run()
    if err != nil {
      c <- fmt.Sprintf("%s", err)
    }
    c <- fmt.Sprintf("%s", "Killed")
  }()

  return c
}


func Term(cd string) <-chan string {
  c := make(chan string)
  go func() {
    comm := cd
    cmd := exec.Command(comm)
    byt, err := cmd.Output()
    if err != nil {
      c <- fmt.Sprintf("%s", err)
    }
    c <- fmt.Sprintf("%s", byt)
  }()

  return c
}

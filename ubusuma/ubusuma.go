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
    var comm string // this is used below

    // Check if there is a space at prefix
    checkfront := strings.HasPrefix(cd, " ")
    if checkfront == true {
      c <- fmt.Sprintf("%s", "Clear space at the Prefix")
      return
    }

    // Check if there is a space at suffix
    checkend := strings.HasSuffix(cd, " ")
    if checkend == true {
      c <- fmt.Sprintf("%s", "Clear space at the Suffix")
      return
    }

    // collect all the spaces and dump them
    // in the spaces slice
    var spaces []int
    for i := 0; i < len(cd); i++ {
      if string(cd[i]) == " " {
        spaces = append(spaces, i)
      }
    }

    // Run the single commant without arguments
    if len(spaces) == 0 {
      comm = cd
      cmd := exec.Command(comm)
      byt, err := cmd.Output()
      if err != nil {
        c <- fmt.Sprintf("%s", err)
      }
      c <- fmt.Sprintf("%s", byt)
      return
    }else if len(spaces) == 1 { // Now deal with one argument
      cmdSlice := spaces[0]
      comm = cd[:cmdSlice] // remember this variable?

      // argument holder
      var args []string

      // holds the argument with a space at the prefix
      // waiting to be removed
      ar := cd[cmdSlice:]
      args = append(args, ar[1:])

      cmd := exec.Command(comm, args...)
      byt, err := cmd.Output()
      if err != nil {
        c <- fmt.Sprintf("%s", err)
      }
      c <- fmt.Sprintf("%s", byt)
      return
    }

  }()

  return c
}

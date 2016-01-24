package main


import (
  "log"
  "os"
  "fmt"
  "net/http"
  "html/template"
  "github.com/kampsy/Go/contype"
  "strings"
  "io/ioutil"
  "encoding/json"
  "time"
  "github.com/syamp/ubusuma"
  "github.com/syamp/random"
)

type WebPage struct {
  Title string
  First_Name string
  Message string
  Distributor string
  Description string
  Release string
  Codename string
}

func main() {
  http.HandleFunc("/home", home)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/", static)

  host := "localhost:2016"
  if len(os.Args) > 1 {
    host = strings.Join(os.Args[1:], " ")
  }
  fmt.Printf("https://%s\n", host)
  err := http.ListenAndServeTLS(host, "cert/Kalibu-Tech.crt", "cert/Kalibu-Tech.key", nil)
  if err != nil {
    log.Printf("%v\n", err)
  }
}

// this function reads files form the drive and returns a slice of bytes
func dirReader(f string) []byte {
  cont, err := ioutil.ReadFile(f)
  if err != nil {
    log.Printf("err : %v\n", err)
    panic("problem when reading")
  }
  return cont
}

// user account type
type UsrStr map[string]string

func rootUsr(cookie string) ([]string, error) {
  jsn := dirReader("usr/root.json")
  var rootJsn UsrStr
  err := json.Unmarshal(jsn, &rootJsn)
  if err != nil {
    return nil, err
  }

  var usr_cookie []string
  usr_cookie = append(usr_cookie, rootJsn["First_Name"])
  return usr_cookie, nil
}

// Returns a string cookie
func rootCoo(cookie string) (string, error) {
  jsn := dirReader("usr/root.json")
  var rootJsn UsrStr
  err := json.Unmarshal(jsn, &rootJsn)
  if err != nil {
    return random.RandStr(5), err // what are olds?
  }

  kie := fmt.Sprintf("%s", rootJsn["Cookie_Key"])
  return kie, nil
}


// This reads Template html files, creates the page and then writes
// the writes to the response writer
func Build(w http.ResponseWriter, p WebPage, tmpFiles []string) {
  tmp := template.New("template")

  // header part of the WebPage
  reVtop, err := ioutil.ReadFile("reVres/tmp/top.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVtop string = string(reVtop)
  tmp.New("reVtop").Parse(str_reVtop)

  // Window material part of the WebPage
  reVmaterial, err := ioutil.ReadFile(fmt.Sprintf("reVres/tmp/%s.html", tmpFiles[0]))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVmaterial string = string(reVmaterial)
  tmp.New("reVmaterial").Parse(str_reVmaterial)


  // body part of the web page
  reVbody, err := ioutil.ReadFile(fmt.Sprintf("reVres/tmp/%s.html", tmpFiles[1]))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVbody string = string(reVbody)
  tmp.New("reVbody").Parse(str_reVbody)

  // bottom part of the web page.
  reVbot, err := ioutil.ReadFile("reVres/tmp/bottom.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVbot string = string(reVbot)
  tmp.New("reVbot").Parse(str_reVbot)


  tmp.Lookup("reVbody").Execute(w, p)
}

// The home page of syamp
func home(w http.ResponseWriter, r *http.Request)  {
  log.Printf("%s: %s \n", r.Method, r.URL.Path)
  cookie, err := r.Cookie("syamp")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }

  // Checks for cookie injections
  cook, err := rootCoo(cookie.Value)
  if err != nil {
    log.Fatal(err)
  }

  if cookie.Value != cook {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }

  // Use the contype to get the write media type from
  // Url
  cont := contype.FileType(r.URL.Path)
  var page WebPage
  page.Title = "Home"


  // Add the name of root user to Webpage struct
  yugi, err := rootUsr(cookie.Value)
  if err != nil {
    log.Fatal(err)
  }
  page.First_Name = yugi[0]

  // Metal returns s reciver Channel of type []slice with value to add to
  // the webpage struct
  metalOut := ubusuma.Metal()
  metalVal := <-metalOut
  page.Distributor = metalVal[0]
  page.Description = metalVal[1]
  page.Release = metalVal[2]
  page.Codename = metalVal[3]

  switch r.Method {
    case "GET":

      query := r.FormValue("stdout")
      if query == "std" {
        // RunningUser returns s reciver Channel of type string.
        running := ubusuma.RunningUser()
        fmt.Fprintf(w, <-running)
        return
      }

      // RunningUser returns s reciver Channel of type string.
      pid := r.FormValue("term")
      if pid != "" {
        ter_msg := ubusuma.Kill(pid)
        fmt.Fprintf(w, <-ter_msg)
        return
      }

      // RunningUser returns s reciver Channel of type string.
      con_cmd := r.FormValue("cmd")
      if con_cmd != "" {
        console := ubusuma.Term(con_cmd)
        fmt.Fprintf(w, <-console)
        return
      }


      w.Header().Set("Content-Type", cont)

      slice := []string {
        "home-window-material",
        "home-body",
      }
      Build(w, page, slice)
    case "POST":
      // This does nothing
      fmt.Fprintf(w, "post home")
  }
}

// This reads Template html files, creates the page and then writes
// the writes to the response writer
func reVtmp(w http.ResponseWriter, p WebPage, body string) {
  tmp := template.New("template")

  reVtop, err := ioutil.ReadFile("reVres/tmp/top.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVtop string = string(reVtop)
  tmp.New("reVtop").Parse(str_reVtop)

  reVbody, err := ioutil.ReadFile(body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVbody string = string(reVbody)
  tmp.New("reVbody").Parse(str_reVbody)

  reVbot, err := ioutil.ReadFile("reVres/tmp/bottom.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var str_reVbot string = string(reVbot)
  tmp.New("reVbot").Parse(str_reVbot)

  tmp.Lookup("reVbody").Execute(w, p)
}

// Returns a slice of strings user data
func queryUser() ([]string, error) {
  jsn := dirReader("usr/root.json")// remember dirReader from line 46
  var rootJsn UsrStr// remember the UsrStr type on line 56?
  err := json.Unmarshal(jsn, &rootJsn)
  if err != nil {
    return nil, err
  }

  var usr_data []string
  usr_data = append(usr_data, rootJsn["First_Name"])
  usr_data = append(usr_data, rootJsn["Last_Name"])
  usr_data = append(usr_data, rootJsn["Password"])
  usr_data = append(usr_data, rootJsn["Cookie_Key"])
  usr_data = append(usr_data, rootJsn["Access"])
  return usr_data, nil
}

// Login page
func login(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s: %s \n", r.Method, r.URL.Path)
  var page WebPage
  page.Title = "Login"
  cont := contype.FileType(r.URL.Path)
  w.Header().Set("Content-Type", cont)
  w.Header().Set("Server", "Syamp")

  switch r.Method {
    case "GET":
      reVtmp(w, page, "reVres/tmp/login-body.html")
    case "POST":
      xusr := r.FormValue("syamp_name")
      xpas := r.FormValue("syamp_pass")
      var check bool
      acc, err := queryUser()
      if err != nil {
        log.Fatal(err)
      }
      if xusr == string(acc[0]) && xpas == string(acc[2]) {
        check = true
      }
      if check == true {
        expiration := time.Now().Add(360 * 24 * time.Hour)
        snack := string(acc[3])
        cookie := http.Cookie{Name: "syamp", Value: snack, Expires: expiration}
        http.SetCookie(w, &cookie)

        http.Redirect(w, r, "/home", http.StatusFound)
        return
      }else {
        page.Message = "syam{p} thinks the information was wrong"
        reVtmp(w, page, "reVres/tmp/login-body.html")
      }
  }

}


//////--logout page---//////
func logout(w http.ResponseWriter, r *http.Request)  {
  log.Printf("%s: %s \n", r.Method, r.URL.Path)
  _, err := r.Cookie("syamp")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }
  expiration := time.Unix(1, 0)
  cookie := http.Cookie{Name: "syamp", MaxAge: -1, Expires: expiration}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/home", http.StatusFound)
}


// config map
type Config map[string]string

func static(w http.ResponseWriter, r *http.Request) {
  jsn := dirReader("etc/config.json")
  var cfig Config
  err := json.Unmarshal(jsn, &cfig)
  if err != nil {
    log.Printf("error: %v\n", err)
  }

  check := strings.Contains(r.URL.Path, cfig["dir"])
  if check == false {
    http.Redirect(w, r, "/home", http.StatusFound)
    return
  }

  path := r.URL.Path[1:]
  data := dirReader(path)
  cont := contype.FileType(r.URL.Path)
  w.Header().Set("Content-Type", cont)
  w.Write(data)
}

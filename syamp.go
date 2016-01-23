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
  "github.com/kalibu_tech_project/ubusuma"
)

type WebPage struct {
  Title string
  First_Name string
  Message string
}

func main() {
  http.HandleFunc("/home", home)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/view", view)
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

// Read json from drive
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
  usr_cookie = append(usr_cookie, rootJsn["Cookie_Key"])
  return usr_cookie, nil
}


/////////----for home-----//////
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

// Home url
func home(w http.ResponseWriter, r *http.Request)  {
  log.Printf("%s: %s \n", r.Method, r.URL.Path)
  cookie, err := r.Cookie("syamp")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }

  cont := contype.FileType(r.URL.Path)
  var page WebPage
  page.Title = "Home"

  // Get the name of the root user
  _, err = rootUsr(cookie.Value)
  if err != nil {
    log.Fatal(err)
  }
  page.First_Name = "gopher"

  switch r.Method {
    case "GET":
      w.Header().Set("Content-Type", cont)
      slice := []string {
        "home-window-material",
        "home-body",
      }
      Build(w, page, slice)
    case "POST":
      fmt.Fprintf(w, "post home")
  }
}


// view full information about the app
func view(w http.ResponseWriter, r *http.Request)  {
  log.Printf("%s: %s \n", r.Method, r.URL.Path)
  cookie, err := r.Cookie("syamp")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }

  cont := contype.FileType(r.URL.Path)
  var page WebPage
  page.Title = "View"

  // Get the name of the root user
  _, err = rootUsr(cookie.Value)
  if err != nil {
    log.Fatal(err)
  }
  page.First_Name = "gopher"

  switch r.Method {
    case "GET":
      webtest := ubusuma.Webapp()
      query := r.FormValue("stdout")
      if query == "std" {
        fmt.Fprintf(w, fmt.Sprintf("%s", <-webtest))
        return
      }

      w.Header().Set("Content-Type", cont)
      slice := []string {
        "home-window-material",
        "view-body",
      }
      Build(w, page, slice)
    case "POST":
      fmt.Fprintf(w, "post home")
  }
}


////////---for log in---///////
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
/////

func queryUser() ([]string, error) {
  jsn := dirReader("usr/root.json")
  var rootJsn UsrStr
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

// Login url
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
        snack := "xxxxx"
        cookie := http.Cookie{Name: "syamp", Value: snack, Expires: expiration}
        http.SetCookie(w, &cookie)

        http.Redirect(w, r, "/home", http.StatusFound)
        return
      }else {
        page.Message = "Error you typed in the wrong Id or Password"
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

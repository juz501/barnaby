package main

import (  
  "html/template"
  "log"
  "net/http"
  "strings"
)

func main() {
  
  http.HandleFunc("/", listen);
  http.HandleFunc("/robots.txt", robots)  
  http.Handle("/static/", noDirListing(http.StripPrefix("/static/", http.FileServer(http.Dir("assets")))))
  log.Fatal(http.ListenAndServe(":80", nil))
}

func noDirListing(h http.Handler) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
    if strings.HasSuffix(r.URL.Path, "/") {
      http.NotFound(w, r)
      return
    }
    h.ServeHTTP(w, r)
  })
}

func robots(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "assets/robots.txt")
}

func listen(w http.ResponseWriter, r *http.Request) {    
  check := func(err error) {
    if err != nil {
      log.Fatal(err)
    }
  }
  defaultName := "Barnaby"  
  queryName := r.URL.Query().Get("name")
  name := defaultName
  if len(queryName) != 0 {
    name = queryName
  }
  
  data := struct {
    Name string
  }{
    Name: name,
  }
  
  t, err := template.ParseFiles("template.tmpl")  
  check(err)
  err = t.Execute(w, data)  
  check(err)  
}